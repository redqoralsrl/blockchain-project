create table if not exists error_log
(
    id            serial primary key,
    timestamp     timestamp,
    ip_address    varchar(255),
    user_agent    varchar(255),
    path          varchar(255),
    http_method   varchar(255),
    requested_url text,
    error_code    integer,
    error_message text,
    stack_trace   text
);

create table if not exists notification
(
    id             SERIAL primary key,

    timestamp      integer,
    transaction_id integer default null,
    log_id         integer default null,

    wallet         VARCHAR(255), -- 지갑 주소의 알림 읽은내역 (트리거로 생성시 from to 2개가 생성)
    is_read        BOOLEAN default false
);

/**
  Blockchain Tracking
*/
create table if not exists tracking_info
(
    id           SERIAL primary key,
    created_at   timestamp default current_timestamp,
    chain_id     integer unique,
    block_height numeric,
    is_operation BOOLEAN   default true
);

create table if not exists crypto_currency
(
    id            SERIAL primary key,
    timestamp     integer,
    morning_date  integer,
    midnight_date integer,
    last_updated  integer,
    eth_price     NUMERIC,
    mmt_price     NUMERIC,
    gmmt_price    NUMERIC,
    matic_price   NUMERIC,
    bnb_price     NUMERIC
);

create table if not exists wallet
(
    id             SERIAL primary key,

    address        VARCHAR(255) unique,
    nick_name      VARCHAR(255) default null,
    profile        VARCHAR(255) default null,
    nonce          numeric      default 0,
    sign_hash      TEXT         default null,
    sign_timestamp integer      default null
);

/*
    block <-> transaction 1 대 다
*/
create table if not exists block
(
    id                SERIAL primary key,
    chain_id          integer,

    difficulty        VARCHAR(255),
    extra_data        TEXT,
    gas_limit         VARCHAR(255),
    gas_used          VARCHAR(255),
    hash              VARCHAR(255),
    logs_bloom        TEXT,
    miner             VARCHAR(42),
    mix_hash          TEXT,
    nonce             VARCHAR(255),
    number_hex        VARCHAR(255),
    number            numeric,
    parent_hash       VARCHAR(255),
    receipts_root     VARCHAR(255),
    sha3_uncles       VARCHAR(255),
    size              VARCHAR(255),
    state_root        VARCHAR(255),
    timestamp         integer,
    total_difficulty  VARCHAR(255),
    transactions_root VARCHAR(255),
    uncles            TEXT,

    unique (hash, chain_id)
);
/*
    transaction <-> contract 1 대 1
    trnasaction <-> log 1 대 다
*/
create table if not exists transaction
(
    id                  SERIAL primary key,
    chain_id            integer,

    block_id            integer references block (id),
    contract_id         integer      default null,

    timestamp           integer,
    block_hash          VARCHAR(255),
    block_number_hex    VARCHAR(255),
    block_number        numeric,
    from_address        VARCHAR(42),
    gas                 VARCHAR(255),
    gas_price           VARCHAR(255),
    hash                VARCHAR(255),
    input               TEXT,
    nonce               VARCHAR(255),
    r                   VARCHAR(255),
    s                   VARCHAR(255),
    to_address          VARCHAR(42)  default null,
    transaction_index   VARCHAR(255),
    type                VARCHAR(255),
    v                   VARCHAR(255),
    value               numeric,
    contract_address    VARCHAR(255) default null,
    cumulative_gas_used VARCHAR(255),
    gas_used            VARCHAR(255),
    logs_bloom          TEXT,
    status              VARCHAR(255),

    unique (hash, chain_id)
);

create table if not exists contract
(
    id                    SERIAL primary key,
    chain_id              integer,

    transaction_create_id integer      default null,

    timestamp             integer,
    hash                  VARCHAR(255),
    name                  TEXT         default null,
    symbol                TEXT         default null,
    decimals              SMALLINT,
    is_erc_20             BOOLEAN      default false,
    is_erc_721            BOOLEAN      default false,
    is_erc_1155           BOOLEAN      default false,
    total_supply          numeric      default 0,
    creator               VARCHAR(42)  default null,

    aws_logo_image        VARCHAR(255) default null,
    aws_banner_image      VARCHAR(255) default null,
    twitter               VARCHAR(255) default null,
    instagram             VARCHAR(255) default null,
    homepage              VARCHAR(255) default null,

    day_volume            numeric      default 0,
    week_volume           numeric      default 0,
    month_volume          numeric      default 0,
    all_volume            numeric      default 0,

    unique (hash, chain_id)
);

alter table transaction add constraint fk_transaction_contract foreign key (contract_id) references contract (id);

alter table contract add constraint fk_contract_transaction foreign key (transaction_create_id) references transaction (id);

alter table contract add column description text default null;
--
-- -- 새로운 열을 추가합니다.
-- ALTER TABLE contract
--     ADD COLUMN new_name TEXT DEFAULT NULL,
--     ADD COLUMN new_symbol TEXT DEFAULT NULL;
--
-- -- 기존 데이터를 새로운 열에 복사합니다.
-- UPDATE contract
-- SET new_name = name,
--     new_symbol = symbol;
--
-- -- 기존 열을 삭제합니다.
-- ALTER TABLE contract
--     DROP COLUMN name,
--     DROP COLUMN symbol;
--
-- -- 새로운 열의 이름을 변경합니다.
-- ALTER TABLE contract
--     RENAME COLUMN new_name TO name;
-- ALTER TABLE contract
--     RENAME COLUMN new_symbol TO symbol;


/*
    log <-> nft_volume 1 대 다
*/
create table if not exists log
(
    id                        SERIAL primary key,
    chain_id                  integer,

    transaction_id            integer references transaction (id),

    address                   VARCHAR(255) default null,
    block_hash                VARCHAR(255),
    block_number_hex          VARCHAR(255),
    block_number              numeric,
    data                      TEXT         default null,
    log_index                 TEXT,
    removed                   BOOLEAN      default false,
    transaction_hash          VARCHAR(255),
    transaction_index         VARCHAR(255),
    timestamp                 integer,
    function                  VARCHAR(255) default null,
    type                      integer      default null,
    dapp                      VARCHAR(255) default null,
    from_address              VARCHAR(42)  default null,
    to_address                VARCHAR(42)  default null,
    value                     numeric      default null,
    token_id                  numeric      default null,
    url                       TEXT         default null,
    name                      TEXT         default null,
    symbol                    TEXT         default null,
    decimals                  integer      default null,
    -- seller            VARCHAR(255) default null,
    -- buyer             VARCHAR(255) default null,
    -- volume            numeric      default null,
    -- volume_symbol     VARCHAR(255) default null,
    -- volume_contract   VARCHAR(255) default null,
    trade_nft_volume          numeric      default null,
    trade_nft_volume_symbol   TEXT         default null,
    trade_nft_volume_contract VARCHAR(255) default null
);
alter table log add column topics VARCHAR(255)[] default null;
alter table log add column erc1155_value numeric[] default null;
alter table log add column erc1155_token_id numeric[] default null;
alter table log add column erc1155_url VARCHAR(255)[] default null;

create table if not exists market_info
(
    id                      SERIAL primary key,
    chain_id                integer,

    log_id                  integer references log (id),

    transaction_hash        VARCHAR(255),
    collection              VARCHAR(255),
    seller                  VARCHAR(255),
    buyer                   VARCHAR(255),
    volume                  numeric,
    volume_symbol           TEXT,
    volume_contract_address TEXT
);

create table if not exists nft_volume
(
    id               SERIAL primary key,
    chain_id         integer,

    log_id           integer references log (id),

    from_address     VARCHAR(42),
    to_address       VARCHAR(42),
    value            numeric default null,
    contract         VARCHAR(255),
    symbol           TEXT,
    timestamp        integer,
    transaction_hash VARCHAR(255),
    event            VARCHAR(255)
);

create table if not exists erc721
(
    id                    SERIAL primary key,
    chain_id              integer,

    contract_id           integer references contract (id),
    wallet_id             integer references wallet (id), -- owner

    token_id              numeric,
    amount                numeric      default 0,
    url                   TEXT         default null,
    image_url             TEXT         default null,
    name                  TEXT         default null,
    description           TEXT         default null,
    external_url          TEXT         default null,
    is_undefined_metadata BOOLEAN      default false,
    aws_image_url         VARCHAR(255) default null,

    unique (contract_id, token_id)
);

create table if not exists erc1155
(
    id                    SERIAL primary key,
    chain_id              integer,

    contract_id           integer references contract (id),

    token_id              numeric,
    amount                numeric      default 0,
    url                   TEXT         default null,
    image_url             TEXT         default null,
    name                  TEXT         default null,
    description           TEXT         default null,
    external_url          TEXT         default null,
    is_undefined_metadata BOOLEAN      default false,
    aws_image_url         VARCHAR(255) default null,

    unique (contract_id, token_id)
);

create table if not exists erc1155_owner
(
    id         SERIAL primary key,
    chain_id   integer,

    erc1155_id integer references erc1155 (id),
    wallet_id  integer references wallet (id),
    amount     numeric default 0,

    unique (erc1155_id, wallet_id)
);

create table if not exists attributes
(
    id           SERIAL primary key,
    chain_id     integer,

    contract_id  integer references contract (id),
    erc721_id    integer references erc721 (id)  null,
    erc1155_id   integer references erc1155 (id) null,

    trait_type   TEXT                            null,
    value        TEXT                            null,
    display_type TEXT                            null
);

create table if not exists market
(
    id                  SERIAL primary key,
    created_at          timestamp    default current_timestamp,
    updated_at          timestamp    default null,

    erc721_id           integer references erc721 (id) null,
    erc1155_id          integer references erc1155 (id) null,
    wallet_id           integer references wallet (id),

    is_dealer           BOOLEAN,                     -- true 판매용 false 제안용
    edem_signer         VARCHAR(255),
    collection          VARCHAR(255),
    token_id            numeric,
    nft_amount          numeric,
    price               numeric,
    protocol_address    VARCHAR(255),
    trade_token_address VARCHAR(255),
    nonce               numeric,
    start_time          integer,
    end_time            integer,
    v                   VARCHAR(255),
    r                   VARCHAR(255),
    s                   VARCHAR(255),
    is_status           VARCHAR(255) default 'SELL', -- SELL 판매 SOLDOUT 판매완료 CANCEL 취소 EXPIRED 만료 SUGGEST 제안 CANCEL_SUGGEST 제안취소 CANCEL_SELL 판매취소 REJECT 제안 거절
    type                VARCHAR(255) default null,   -- SELL 판매 / SOLDOUT_SELL 판매로 판매완료 / SUGGEST 제안 / SOLDOUT_SUGGEST 제안승낙으로 판매완료 / CANCEL_SELL 판매취소 / CANCEL_SUGGEST 제안 취소 / REJECT_SUGGEST 제안 거절
    sign_hash           VARCHAR(255)
);

-- -- trigger 1
-- create or replace function add_notification_transaction() returns trigger as $$
-- begin
--     insert into notification (timestamp, transaction_id, log_id, wallet, is_read)
--     values (new.timestamp, new.id, null, new.from_address, false);
--     insert into notification (timestamp, transaction_id, log_id, wallet, is_read)
--     values (new.timestamp, new.id, null, new.to_address, false);
--
--     return new;
-- end;
-- $$ language plpgsql;
--
-- create trigger transaction_insert_trigger
--     after insert
--     on transaction
--     for each row
-- execute function add_notification_transaction();
--
-- -- trigger 2
-- create or replace function add_notification_log() returns trigger as $$
-- begin
--     insert into notification (timestamp, transaction_id, log_id, wallet, is_read)
--     values (new.timestamp, null, new.id, new.from_address, false);
--     insert into notification (timestamp, transaction_id, log_id, wallet, is_read)
--     values (new.timestamp, null, new.id, new.to_address, false);
--
--     return new;
-- end;
-- $$ language plpgsql;
--
-- create trigger log_insert_trigger
--     after insert
--     on log
--     for each row
-- execute function add_notification_log();

-- TODO: chain_id
-- initial setting
insert into tracking_info (chain_id, block_height)
values (8989, 2);
insert into tracking_info (chain_id, block_height)
values (898989, 0);
insert into tracking_info (chain_id, block_height, is_operation)
values (1, 0, false);
insert into tracking_info (chain_id, block_height, is_operation)
values (137, 0, false);
insert into tracking_info (chain_id, block_height, is_operation)
values (56, 0, false);

-- TODO: chain_id
-- initial setting smartcontract
insert into contract (chain_id, transaction_create_id, timestamp, hash, name, symbol, decimals, is_erc_20, is_erc_721,
                      total_supply, creator)
values (8989, null, 0, '0x0000000000000000000000000000000000001000', 'validate contract', '', 0, false, false, 0, ''),
       (8989, null, 0, '0x0000000000000000000000000000000000001001', 'slash contract', '', 0, false, false, 0, ''),
       (8989, null, 0, '0x0000000000000000000000000000000000001002', 'system reward contract', '', 0, false, false, 0,
        ''),
       (8989, null, 0, '0x0000000000000000000000000000000000001003', 'light client contract', '', 0, false, false, 0,
        ''),
       (8989, null, 0, '0x0000000000000000000000000000000000001004', 'token hub contract', '', 0, false, false, 0, ''),
       (8989, null, 0, '0x0000000000000000000000000000000000001005', 'relayer incentivize contract', '', 0, false,
        false, 0, ''),
       (8989, null, 0, '0x0000000000000000000000000000000000001006', 'relayer hub contract', '', 0, false, false, 0,
        ''),
       (8989, null, 0, '0x0000000000000000000000000000000000001007', 'gov hub contract', '', 0, false, false, 0, ''),
       (8989, null, 0, '0x0000000000000000000000000000000000001008', 'token manager contract', '', 0, false, false, 0,
        ''),
       (8989, null, 0, '0x0000000000000000000000000000000000002000', 'cross chain contract', '', 0, false, false, 0,
        ''),
       (8989, null, 0, '0x0000000000000000000000000000000000007001', 'staking pool contract', '', 0, false, false, 0,
        ''),
       (8989, null, 0, '0x0000000000000000000000000000000000007002', 'governance contract', '', 0, false, false, 0, ''),
       (8989, null, 0, '0x0000000000000000000000000000000000007003', 'chain config contract', '', 0, false, false, 0,
        ''),
       (8989, null, 0, '0x0000000000000000000000000000000000007004', 'runtime upgrade contract', '', 0, false, false, 0,
        ''),
       (8989, null, 0, '0x0000000000000000000000000000000000007005', 'deployer proxy contract', '', 0, false, false, 0,
        '');

-- TODO: chain_id
insert into block (chain_id, difficulty, extra_data, gas_limit, gas_used, hash, logs_bloom, miner, mix_hash,
                   nonce, number_hex, "number", parent_hash, receipts_root, sha3_uncles, "size", state_root,
                   "timestamp", total_difficulty, transactions_root, uncles)
values (8989, '0x1',
        '0x00000000000000000000000000000000000000000000000000000000000000003ef8cb3c73f0dfa760981d2bcd57a7e9d8535e6f8ee898ba66d6551f80dbca32c81157b02847812955e6d8342150078e969d75396b93918c9ce0cf1bcfcde214610130883e92098cd28705eb065da4d13a1866e17c8807701599fd5bad97ffc63387e8c9be9aac23d8fe2b91860a87a1e38f661c7e73563e266ad1d026f4cedafea84243d4be3cada9ec52d7d1f722c4b60298dfecaa9be8eec7069c59dccf630f0dac5b5c3cc5abe25ebf70f9d0cf7b4ef4cb4d6cb28a9500e65fd02433d4eced9fa7435a4cec7381d1a6c7cef3646fe14fbf86e5f4f390a8d502d4c07a1a7c98c803632333c3410f1cb0fef70c9d8388c7e92dca30a867bec0f83152c372998073cc4920793d66f917dc37c4a3a33d6a31228ee85c6b0c452c9fa3d3d2bda54e55a480f4adb3abdeb96e890ef2da91249ef620c9fc0c4f4978398829485b3968427ba874a7c4451e3be77e980c9719473492073a26416b048e41366d7db378c8c7110a6187f0c475902f785e727606faaf8474adc12a718d2a6e6d4dff4db5a1e124f647beceb64d84d31ff6162d423c38faf3297ef9bc9fad3dac03eeab2783a4b60e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000',
        '0x3b9aca00', '0x0', '0x20e341740663b942a46c1fba95b1329ef0e8516f56cfd639583572d73932e65a',
        '0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000',
        '0x0000000000000000000000000000000000000000',
        '0x0000000000000000000000000000000000000000000000000000000000000000', '0x0000000000000000', '0x0', 0,
        '0x0000000000000000000000000000000000000000000000000000000000000000',
        '0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421',
        '0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347', '0x404',
        '0x775adab72016950b868138c99fab83ff81bb5a641e4f7d3d6e35738f0297ee14', 1670262300, '0x1',
        '0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421', '');
insert into block (chain_id, difficulty, extra_data, gas_limit, gas_used, hash, logs_bloom, miner, mix_hash,
                   nonce, number_hex, "number", parent_hash, receipts_root, sha3_uncles, "size", state_root,
                   "timestamp", total_difficulty, transactions_root, uncles)
values (8989, '0x2',
        '0xd883010108846765746888676f312e31382e31856c696e75780000003a4dc509add08dd903b448a02aef8c8cb0a153fc8eecb8a3da4919f65dfc1e9b38c5b4ca6e441dda549f498b0bcf5aa90bc7ffd7d572ba396b488f23b93ae91be34530bf01',
        '0x3b5f2f37', '0x4f997c', '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51',
        '0x0000000050000084210000400000030000801000000004000010000000000240000010000010200020000000402000040080000000080800800000000000000100000000404000840404000080000000202000000980000000001100000000009090102040000404004244000008000000100020000080101510000000004000000000000800000000400000008080010000044000000000000000020000400200000000000000800004400100000000000000003000c808000080800020020000000000000000000011000000580000004000600000000002004004000000040080020000101000000810000800880000000000010004000200000000800000',
        '0x0f0dac5b5c3cc5abe25ebf70f9d0cf7b4ef4cb4d',
        '0x0000000000000000000000000000000000000000000000000000000000000000', '0x0000000000000000', '0x1', 1,
        '0x20e341740663b942a46c1fba95b1329ef0e8516f56cfd639583572d73932e65a',
        '0x712add76700af76cc3f31ca32213777f2cf41f96b7791c63936c0b8428c1845e',
        '0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347', '0x5cc',
        '0x5a19717aad43ee86717bc29390bdadd6257dc364ae6a4e564070c852ce10b4af', 1672714832, '0x3',
        '0x38f3065eaa9d66451a6effaaa71339ffc2f9ff47c02e7e82eefe05d493963d32', '');

insert into transaction (chain_id, block_id, contract_id, "timestamp", block_hash, block_number_hex,
                         block_number, from_address, gas, gas_price, hash, "input", nonce, r, s, to_address,
                         transaction_index, "type", v, value, contract_address, cumulative_gas_used, gas_used,
                         logs_bloom, status)
values (8989, 2, null, 1672714832, '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x0f0dac5b5c3cc5abe25ebf70f9d0cf7b4ef4cb4d', '0x7fffffffffffffff', '0x0',
        '0xbeb8f95846c07e11f1d1980c58374a1d6af552a9bdb0f8b32fab3cfe24a045e2', '0xe1c7392a', '0x7',
        '0x1039b007b223c20ac49050a9161d24d6996d302c6b0f48cce944cf6f21b482c8',
        '0x3e4098f3f4b5b987c02321a817df1e1e34cea3cb4fc27c9af07b14ae1f4579dd',
        '0x0000000000000000000000000000000000007005', '0x7', '0x0', '0x465d', 0, '', '0x4f997c', '0x35b0d',
        '0x00000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000200000000000000000000000000000000000000000000000000800000000000000000000000200000000800000',
        '0x1'),
       (8989, 2, null, 1672714832, '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x0f0dac5b5c3cc5abe25ebf70f9d0cf7b4ef4cb4d', '0x7fffffffffffffff', '0x0',
        '0xa25c8158381f0add4a26ccbfa5eaccd5f441e4062d51f23aabc666368b06ad84', '0xe1c7392a', '0x4',
        '0x65b69b1f9ff868c3c6d74d9e22bbcdab65b451f52cbbf23a58b52e44d932a5ba',
        '0x58944b39ed9f307f83625560e8c66135c072ec367e5036f809a5d051762a03b3',
        '0x0000000000000000000000000000000000007002', '0x4', '0x0', '0x465d', 0, '', '0x44ac80', '0x31298',
        '0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000040002000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000',
        '0x1'),
       (8989, 2, null, 1672714832, '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x0f0dac5b5c3cc5abe25ebf70f9d0cf7b4ef4cb4d', '0x7fffffffffffffff', '0x0',
        '0xed69f0fc15d2e46606830c67d553228abea2ef009085c6616d12907cc7c2adfe', '0xe1c7392a', '0x3',
        '0x3c0e4dd7366eab6cad3e0ff4e3e99cf37802c5b5f49b2a4f3b4510ba2f1db1da',
        '0x45214b00a9e019b79cf3d7ff9e22e1b6009481d51ea05386095397ff7e553c8c',
        '0x0000000000000000000000000000000000007001', '0x3', '0x0', '0x465e', 0, '', '0x4199e8', '0x2ef9a',
        '0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000',
        '0x1'),
       (8989, 2, null, 1672714832, '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x0f0dac5b5c3cc5abe25ebf70f9d0cf7b4ef4cb4d', '0x7fffffffffffffff', '0x0',
        '0xf76f8c63a1980b61951b1f04ce491e31cde0b8c92b6bffa302e7fa09b9b64ef0', '0xe1c7392a', '0x5',
        '0xf53f21a5f2a011ba0d6dd9b0a52172ea816217ca8e96e7c860501169698f237e',
        '0x742a21e0c3c96c97a190bd0bddc1c905a5bdc337005548712395fac9273fa4fb',
        '0x0000000000000000000000000000000000007003', '0x5', '0x0', '0x465e', 0, '', '0x48f5f5', '0x44975',
        '0x00000000100000040000000000000000008000000000040000000000000000000000000000102000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000001000400000000000040000000000000000200000001010000000000040000000000000000000000000000000800000000000000000000000000200000000000000000000008000004001000000000000000020008000000000000000000000000000000000000000000000480000000000000000000000000000000000000000000000101000000010000000000000000000010000000000000000000000',
        '0x1'),
       (8989, 2, null, 1672714832, '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x0f0dac5b5c3cc5abe25ebf70f9d0cf7b4ef4cb4d', '0x7fffffffffffffff', '0x0',
        '0x2f97fc03e985280ebcdff904ddb078b942b750951858b4e7da22a4d5b0c454ee', '0xe1c7392a', '0x6',
        '0xe8c4f6725c653ac380c87ea531ad2e1fd0dacdf1fc1a65c165fd4aa5cc8037c',
        '0xc6463fbe3bdebc98567a054b5121365705438ee92ce49c43c595d028f7a5213',
        '0x0000000000000000000000000000000000007004', '0x6', '0x0', '0x465e', 0, '', '0x4c3e6f', '0x3487a',
        '0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000',
        '0x1'),
       (8989, 2, null, 1672714832, '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x0f0dac5b5c3cc5abe25ebf70f9d0cf7b4ef4cb4d', '0x7fffffffffffffff', '0x0',
        '0x8306e55b839aa572857f11d8cf44a1c58eff7bc6306fe3326ca22cf9693d346a', '0xe1c7392a', '0x2',
        '0xebab0f80b5ab20f12f7b3c05ade2e9c3cfaf67e0a8418e4de6a64418a29c34a',
        '0x39e8ddb121b4fb827d1d3561a4e7c9d28236e5f76f2150ec1f7fa6462a2e3de5',
        '0x0000000000000000000000000000000000001002', '0x2', '0x0', '0x465e', 0, '', '0x3eaa4e', '0x3bf9c',
        '0x00000000000000000000004000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000080000000000000000000000004000000000000080000000000000000000000000000000000000000000000000000000000000',
        '0x1'),
       (8989, 2, null, 1672714832, '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x0f0dac5b5c3cc5abe25ebf70f9d0cf7b4ef4cb4d', '0x7fffffffffffffff', '0x0',
        '0x7845f28ffa744ac6cf21a0ee852f716d87c9a815691550900248ddcf045a13fe', '0xe1c7392a', '0x1',
        '0x645758b453339e805c26a96cd0fc40d39e919c438c32ea59393893381a31811f',
        '0x1756e71bdfa01a1dcbe399b774d7837de1307fb8154bcef0f46afa0ff2b79a0e',
        '0x0000000000000000000000000000000000001001', '0x1', '0x0', '0x465e', 0, '', '0x3aeab2', '0x2ef42',
        '0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000',
        '0x1'),
       (8989, 2, null, 1672714832, '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x0f0dac5b5c3cc5abe25ebf70f9d0cf7b4ef4cb4d', '0x7fffffffffffffff', '0x0',
        '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0xe1c7392a', '0x0',
        '0xf9d0fa93ecdd9943acf7a3ed36ab16fb58faaa2815b18340bc390f7143e8da42',
        '0x17a0912bd092c1a9329a04c2e48c0314ca2cc2ed15ffd3e39c1dec3d41e04ff0',
        '0x0000000000000000000000000000000000001000', '0x0', '0x0', '0x465e', 0, '', '0x37fb70', '0x37fb70',
        '0x00000000400000802100000000000300000000000000000000100000000002400000000000000000200000004020000400800000000808000000000000000001000000004040008400040000800000002000000009800000000001000000000090900020000004000040400000080000001000000000800001100000000000000000000008000000004000000080000000000440000000000000000000004000000000000000000000040000000000000000000010004808000080000020020000000000000000000011000000000000004000400000000002000004000000040000020000000000000800000000880000000000000004000200000000800000',
        '0x1');

insert into log (chain_id, transaction_id, address, block_hash, block_number_hex, block_number, "data",
                 log_index, removed, transaction_hash, transaction_index, "timestamp", "function", "type", dapp,
                 from_address, to_address, value, token_id, url, name, symbol, decimals,
                 trade_nft_volume, trade_nft_volume_symbol,
                 trade_nft_volume_contract, topics)
values (8989, 1, '0x0000000000000000000000000000000000007005',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1, '0x', '0x1f', false,
        '0xbeb8f95846c07e11f1d1980c58374a1d6af552a9bdb0f8b32fab3cfe24a045e2', '0x7', 1672714832, '', 0, '',
        '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '', '', '', 0,
         0, '', '',
        '{0x861a21548a3ee34d896ccac3668a9d65030aaf2cb7367a2ed13608014016a032,0x0000000000000000000000003ef8cb3c73f0dfa760981d2bcd57a7e9d8535e6f}'),
       (8989, 2, '0x0000000000000000000000000000000000007002',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000007080',
        '0x16', false, '0xa25c8158381f0add4a26ccbfa5eaccd5f441e4062d51f23aabc666368b06ad84', '0x4', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '', '{0x7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828}'),
       (8989, 4, '0x0000000000000000000000000000000000007003',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000015',
        '0x17', false, '0xf76f8c63a1980b61951b1f04ce491e31cde0b8c92b6bffa302e7fa09b9b64ef0', '0x5', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0, 0, '', '', '{0x1c4cfc6dcf4219ed649285020aedf5d064480d1acdf4b8c75b397abd5910f40c}'),
       (8989, 6, '0x0000000000000000000000000000000000001002',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x00000000000000000000000032852293fc76d99b130250252f5a1803f613c1500000000000000000000000000000000000000000000000000000000000002710',
        '0x15', false, '0x8306e55b839aa572857f11d8cf44a1c58eff7bc6306fe3326ca22cf9693d346a', '0x2', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0, 0, '', '', '{0x55f2d69d9dbe97a29594a0106efd4f56ce72ec40b82d3583d53ef836e11bd00f}'),
       (8989, 8, '0x0000000000000000000000000000000000001000',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x0000000000000000000000003ef8cb3c73f0dfa760981d2bcd57a7e9d8535e6f00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        '0x0', false, '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0x0', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '',
        '{0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d,0x0000000000000000000000003ef8cb3c73f0dfa760981d2bcd57a7e9d8535e6f}'),
       (8989, 4, '0x0000000000000000000000000000000000007003',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007080',
        '0x18', false, '0xf76f8c63a1980b61951b1f04ce491e31cde0b8c92b6bffa302e7fa09b9b64ef0', '0x5', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '', '{0x33c8012b0f51f8c1a1e525ea046da837d0eb4fa7473cd863e0bfb73a4f475a5a}'),
       (8989, 8, '0x0000000000000000000000000000000000001000',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x0000000000000000000000008ee898ba66d6551f80dbca32c81157b02847812900000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        '0x1', false, '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0x0', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '',
        '{0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d,0x0000000000000000000000008ee898ba66d6551f80dbca32c81157b028478129}'),
       (8989, 4, '0x0000000000000000000000000000000000007003',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a',
        '0x19', false, '0xf76f8c63a1980b61951b1f04ce491e31cde0b8c92b6bffa302e7fa09b9b64ef0', '0x5', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '', '{0x5aa72ebd12c45515403eef36583106e321b8707946a6ae621f5f393ee0c9677b}'),
       (8989, 8, '0x0000000000000000000000000000000000001000',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x00000000000000000000000055e6d8342150078e969d75396b93918c9ce0cf1b00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        '0x2', false, '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0x0', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0, 0, '', '',
        '{0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d,0x00000000000000000000000055e6d8342150078e969d75396b93918c9ce0cf1b}'),
       (8989, 4, '0x0000000000000000000000000000000000007003',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a',
        '0x1a', false, '0xf76f8c63a1980b61951b1f04ce491e31cde0b8c92b6bffa302e7fa09b9b64ef0', '0x5', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '', '{0x67da1e9c07e7b373ed5e18cc8355caf6dfe18ab4472ec575600a2172772c6204}');
insert into log (chain_id, transaction_id, address, block_hash, block_number_hex, block_number, "data",
                 log_index, removed, transaction_hash, transaction_index, "timestamp", "function", "type", dapp,
                 from_address, to_address, value, token_id, url, name, symbol, decimals,
                trade_nft_volume, trade_nft_volume_symbol,
                 trade_nft_volume_contract, topics)
values (8989, 8, '0x0000000000000000000000000000000000001000',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x000000000000000000000000cfcde214610130883e92098cd28705eb065da4d100000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        '0x3', false, '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0x0', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '',
        '{0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d,0x000000000000000000000000cfcde214610130883e92098cd28705eb065da4d1}'),
       (8989, 4, '0x0000000000000000000000000000000000007003',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001',
        '0x1b', false, '0xf76f8c63a1980b61951b1f04ce491e31cde0b8c92b6bffa302e7fa09b9b64ef0', '0x5', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '', '{0x0a677ce4509bf46fe9bdf65c86abe71921755a78494111b1caa25df328ffcd1c}'),
       (8989, 4, '0x0000000000000000000000000000000000007003',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000',
        '0x1c', false, '0xf76f8c63a1980b61951b1f04ce491e31cde0b8c92b6bffa302e7fa09b9b64ef0', '0x5', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '', '{0xb191e5acbef9e4b8ce0f17af112f8984f92833324657b89fe39768885f81b6ce}'),
       (8989, 4, '0x0000000000000000000000000000000000007003',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a968163f0a57b400000',
        '0x1d', false, '0xf76f8c63a1980b61951b1f04ce491e31cde0b8c92b6bffa302e7fa09b9b64ef0', '0x5', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '', '{0x207082661d623a88e041ad2d52c2d4ddc719880c70c3ab44aa81accff9bd86ed}'),
       (8989, 4, '0x0000000000000000000000000000000000007003',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000de0b6b3a7640000',
        '0x1e', false, '0xf76f8c63a1980b61951b1f04ce491e31cde0b8c92b6bffa302e7fa09b9b64ef0', '0x5', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0, 0, '', '', '{0x973f438cb6bc47d284033b6113687c6087f4fb7a3395b03597578ae1259bf23c}'),
       (8989, 8, '0x0000000000000000000000000000000000001000',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x0000000000000000000000000f0dac5b5c3cc5abe25ebf70f9d0cf7b4ef4cb4d00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        '0x8', false, '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0x0', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '',
        '{0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d,0x0000000000000000000000000f0dac5b5c3cc5abe25ebf70f9d0cf7b4ef4cb4d}'),
       (8989, 8, '0x0000000000000000000000000000000000001000',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x0000000000000000000000006cb28a9500e65fd02433d4eced9fa7435a4cec7300000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        '0x9', false, '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0x0', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0, 0, '', '',
        '{0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d,0x0000000000000000000000006cb28a9500e65fd02433d4eced9fa7435a4cec73}'),
       (8989, 8, '0x0000000000000000000000000000000000001000',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x00000000000000000000000081d1a6c7cef3646fe14fbf86e5f4f390a8d502d400000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        '0xa', false, '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0x0', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0, 0, '', '',
        '{0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d,0x00000000000000000000000081d1a6c7cef3646fe14fbf86e5f4f390a8d502d4}'),
       (8989, 8, '0x0000000000000000000000000000000000001000',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x000000000000000000000000c07a1a7c98c803632333c3410f1cb0fef70c9d8300000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        '0xb', false, '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0x0', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '',
        '{0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d,0x000000000000000000000000c07a1a7c98c803632333c3410f1cb0fef70c9d83}'),
       (8989, 8, '0x0000000000000000000000000000000000001000',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x00000000000000000000000088c7e92dca30a867bec0f83152c372998073cc4900000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        '0xc', false, '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0x0', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '',
        '{0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d,0x00000000000000000000000088c7e92dca30a867bec0f83152c372998073cc49}');
insert into log (chain_id, transaction_id, address, block_hash, block_number_hex, block_number, "data",
                 log_index, removed, transaction_hash, transaction_index, "timestamp", "function", "type", dapp,
                 from_address, to_address, value, token_id, url, name, symbol, decimals,
                  trade_nft_volume, trade_nft_volume_symbol,
                 trade_nft_volume_contract, topics)
values (8989, 8, '0x0000000000000000000000000000000000001000',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x00000000000000000000000020793d66f917dc37c4a3a33d6a31228ee85c6b0c00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        '0xd', false, '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0x0', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '',
        '{0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d,0x00000000000000000000000020793d66f917dc37c4a3a33d6a31228ee85c6b0c}'),
       (8989, 8, '0x0000000000000000000000000000000000001000',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x0000000000000000000000003a1866e17c8807701599fd5bad97ffc63387e8c900000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        '0x4', false, '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0x0', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '',
        '{0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d,0x0000000000000000000000003a1866e17c8807701599fd5bad97ffc63387e8c9}'),
       (8989, 8, '0x0000000000000000000000000000000000001000',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x000000000000000000000000be9aac23d8fe2b91860a87a1e38f661c7e73563e00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        '0x5', false, '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0x0', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '',
        '{0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d,0x000000000000000000000000be9aac23d8fe2b91860a87a1e38f661c7e73563e}'),
       (8989, 8, '0x0000000000000000000000000000000000001000',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x000000000000000000000000266ad1d026f4cedafea84243d4be3cada9ec52d700000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        '0x6', false, '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0x0', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '',
        '{0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d,0x000000000000000000000000266ad1d026f4cedafea84243d4be3cada9ec52d7}'),
       (8989, 8, '0x0000000000000000000000000000000000001000',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x000000000000000000000000d1f722c4b60298dfecaa9be8eec7069c59dccf6300000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        '0x7', false, '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0x0', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0, 0, '', '',
        '{0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d,0x000000000000000000000000d1f722c4b60298dfecaa9be8eec7069c59dccf63}'),
       (8989, 8, '0x0000000000000000000000000000000000001000',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x000000000000000000000000452c9fa3d3d2bda54e55a480f4adb3abdeb96e8900000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        '0xe', false, '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0x0', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '',
        '{0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d,0x000000000000000000000000452c9fa3d3d2bda54e55a480f4adb3abdeb96e89}'),
       (8989, 8, '0x0000000000000000000000000000000000001000',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x0000000000000000000000000ef2da91249ef620c9fc0c4f4978398829485b3900000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        '0xf', false, '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0x0', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '',
        '{0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d,0x0000000000000000000000000ef2da91249ef620c9fc0c4f4978398829485b39}'),
       (8989, 8, '0x0000000000000000000000000000000000001000',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x00000000000000000000000068427ba874a7c4451e3be77e980c97194734920700000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        '0x10', false, '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0x0', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '',
        '{0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d,0x00000000000000000000000068427ba874a7c4451e3be77e980c971947349207}'),
       (8989, 8, '0x0000000000000000000000000000000000001000',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x0000000000000000000000003a26416b048e41366d7db378c8c7110a6187f0c400000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        '0x11', false, '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0x0', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0, 0, '', '',
        '{0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d,0x0000000000000000000000003a26416b048e41366d7db378c8c7110a6187f0c4}'),
       (8989, 8, '0x0000000000000000000000000000000000001000',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x00000000000000000000000075902f785e727606faaf8474adc12a718d2a6e6d00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        '0x12', false, '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0x0', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '',
        '{0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d,0x00000000000000000000000075902f785e727606faaf8474adc12a718d2a6e6d}');
insert into log (chain_id, transaction_id, address, block_hash, block_number_hex, block_number, "data",
                 log_index, removed, transaction_hash, transaction_index, "timestamp", "function", "type", dapp,
                 from_address, to_address, value, token_id, url, name, symbol, decimals,
                 trade_nft_volume, trade_nft_volume_symbol,
                 trade_nft_volume_contract, topics)
values (8989, 8, '0x0000000000000000000000000000000000001000',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x0000000000000000000000004dff4db5a1e124f647beceb64d84d31ff6162d4200000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        '0x13', false, '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0x0', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0, 0, '', '',
        '{0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d,0x0000000000000000000000004dff4db5a1e124f647beceb64d84d31ff6162d42}'),
       (8989, 8, '0x0000000000000000000000000000000000001000',
        '0x772484226acff59f9996ba52060aab50bb5c865a03f79038abc89814622a1b51', '0x1', 1,
        '0x0000000000000000000000003c38faf3297ef9bc9fad3dac03eeab2783a4b60e00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        '0x14', false, '0x3745b5eb9b1fe9b74011dc64a768ec8ad2fd52840d63642fd8699c29022c89d3', '0x0', 1672714832, '', 0,
        '', '0x0000000000000000000000000000000000000000', '0x0000000000000000000000000000000000000000', 0, 0, '',
        '', '', 0,  0, '', '',
        '{0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d,0x0000000000000000000000003c38faf3297ef9bc9fad3dac03eeab2783a4b60e}');

-------------------------- chain info
create table if not exists chain
(
    id               SERIAL primary key,
    chain_id         integer,
    contract_address varchar(255) default '0x',
    name             text,
    symbol           text,
    decimals         integer not null,
    image_url        text,
    site             text,
    scan_site        text,

    unique (chain_id, contract_address)
);

-- index setting
create index idx_wallet_address on wallet (lower(address));

create index idx_chain_chain_id_contract_address on chain (chain_id, contract_address);
create index idx_chain_chain_id_contract_address_lower on chain (chain_id, lower(contract_address));
create index idx_chain_chain_id on chain (chain_id);

create index idx_contract_hash_chain_id on contract (lower(hash), chain_id);
create index idx_contract_id ON contract(id);
create index idx_contract_chain_id ON contract(chain_id);
create index idx_contract_hash ON contract(lower(hash));
create index idx_contract_transaction_create_id ON contract(transaction_create_id);

create index idx_erc721_contract_id_wallet_id on erc721 (contract_id, wallet_id);
create index idx_erc721_contract_id_chain_id_token_id ON erc721 (contract_id, chain_id, token_id);
create index idx_erc721_wallet_id ON erc721 (wallet_id);
create index idx_erc721_metadata_image_url ON erc721 (is_undefined_metadata, image_url);
create index idx_erc721_contract_chain_token ON erc721 (contract_id, chain_id, token_id);

create index idx_erc1155_contract_id_wallet_id on erc1155 (contract_id);
create index idx_erc1155_id_contract ON erc1155 (id, contract_id);

create index idx_erc1155_owner_wallet_erc1155 ON erc1155_owner (wallet_id, erc1155_id);

create index idx_log_chain_id ON log(chain_id);
create index idx_log_address ON log(lower(address));
create index idx_log_from_address ON log(lower(from_address));
create index idx_log_to_address ON log(lower(to_address));
create index idx_log_transaction_id ON log(transaction_id);
create index idx_log_transaction_hash ON log(transaction_hash);
create index idx_log_timestamp ON log(timestamp);

create index idx_transaction_id on transaction (id);
create index idx_transaction_chain_id ON transaction(chain_id);
create index idx_transaction_from_address ON transaction(lower(from_address));
create index idx_transaction_to_address ON transaction(lower(to_address));
create index idx_transaction_hash ON transaction(lower(hash));
create index idx_transaction_timestamp ON transaction(timestamp);

-- TODO: chain info
insert into chain (chain_id, contract_address, name, symbol, decimals, image_url, site, scan_site)
values (8989, '0x', 'Giant Mammoth', 'GMMT', 18, 'https://d229hptfcazewb.cloudfront.net/coin/gmmt.svg',
        'https://gmmtchain.io', 'https://scan.gmmtchain.io/'),
       (898989, '0x', 'Giant Mammoth', 'GMMT', 18, 'https://d229hptfcazewb.cloudfront.net/coin/gmmt.svg',
        'https://gmmtchain.io', 'https://testnet-scan.gmmtchain.io/'),
       (8989, '0x6fA62D97C4dd43eeA64c7512D4F3501E1170413F', 'Marzmine', 'MZM', 18,
        'https://d229hptfcazewb.cloudfront.net/coin/mzm.svg', 'https://gmmtchain.io', 'https://scan.gmmtchain.io/'),
       (8989, '0x734bf3a059ee840a910e02e477049ef9c1a644ab', 'Cockfight Network', 'CFN', 18,
        'https://d229hptfcazewb.cloudfront.net/coin/cfn.svg', 'https://gmmtchain.io', 'https://scan.gmmtchain.io/'),
       (8989, '0xddfe042e414978e59dd16c8ac8487160ebaf24a5', 'Wrapped Giant Mammoth', 'WGMMT', 18,
        'https://d229hptfcazewb.cloudfront.net/coin/wgmmt.svg', 'https://gmmtchain.io', 'https://scan.gmmtchain.io/'),
       (8989, '0xed80bde82865f262c9bd6d295782f5b0172b19bb', 'WalkDoni', 'WDT', 18,
        'https://d229hptfcazewb.cloudfront.net/coin/wdt.svg', 'https://gmmtchain.io', 'https://scan.gmmtchain.io/'),
       (8989, '0x1f1d3fb5d16171cc4586ebc890ef6a11e8951296', 'Ai Bitcoin Pick', 'ABP', 18,
        'https://d229hptfcazewb.cloudfront.net/coin/abp.svg', 'https://gmmtchain.io', 'https://scan.gmmtchain.io/'),
       (8989, '0x48a590907f15f54137b352037ba09e1638b592ac', 'IvoryToken', 'IVY', 18,
        'https://d229hptfcazewb.cloudfront.net/coin/ivy.svg', 'https://gmmtchain.io', 'https://scan.gmmtchain.io/'),
       (8989, '0x7BDdff151D5AeE4FfDA2881f64F47B585855013F', 'Espero', 'ESP', 18,
        'https://d229hptfcazewb.cloudfront.net/coin/esp.svg', 'https://gmmtchain.io', 'https://scan.gmmtchain.io/');

-- (1, '0x', 'Ethereum', 'ETH', 18, 'https://d229hptfcazewb.cloudfront.net/coin/eth.svg', 'https://ethereum.org/en',
--  'https://etherscan.io/'),
-- (137, '0x', 'Polygon', 'MATIC', 18, 'https://d229hptfcazewb.cloudfront.net/coin/matic.svg',
--  'https://polygon.technology', 'https://polygonscan.com/'),
-- (56, '0x', 'BNB Chain', 'BNB', 18, 'https://d229hptfcazewb.cloudfront.net/coin/bnb.svg',
--  'https://www.bnbchain.org/en', 'https://bscscan.com/'),