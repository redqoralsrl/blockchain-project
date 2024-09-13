#!/bin/bash
if ps aux | grep '[.]/main'; then
        echo ""
else
        echo "gmwallet socket 프로세스가 실행되지 않았습니다. tmux 세션을 재시작합니다."
        # 기존에 실행 중이던 tmux 세션 종료
        tmux kill-session -t gmwallet 2>/dev/null
        # 백그라운드에서 새로운 tmux 세션 생성
        tmux new -d -s gmwallet
        # 생성된 세션에서 빌드된 프로그램 실행
        tmux send-keys -t gmwallet './main' C-m
fi
while true
do
    if ps aux | grep '[.]/main'; then
            echo ""
    else
            echo "gmwallet socket 프로세스가 실행되지 않았습니다. tmux 세션을 재시작합니다."
            # 기존에 실행 중이던 tmux 세션 종료
            tmux kill-session -t gmwallet 2>/dev/null
            # 백그라운드에서 새로운 tmux 세션 생성
            tmux new -d -s gmwallet
            # 생성된 세션에서 빌드된 프로그램 실행
            tmux send-keys -t gmwallet './main' C-m
    fi
    sleep 60
done