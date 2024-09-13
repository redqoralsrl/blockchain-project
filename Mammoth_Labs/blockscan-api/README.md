# Blockscan-go

---

### Branch
> main
> 
> Production Echo API Service / No Cronjob / No Rpc Server

> develop
> 
> Not planning

> track
> 
> Tracking Service / Cronjob / Rpc Server

---

## Tracking Service 인프라

Singapore EC2 / RDS

## Echo Service

Production - Singapore ECS / RDS

Develop - Not yet

---

prettyJSON, _ := json.MarshalIndent(results, "", "    ")
fmt.Println("results : ", string(prettyJSON))