cd ./src/github.com/karux/testRESTService
sls invoke --function hello --path ./data/data.json # -log 
sls invoke --function world --path ./data/data.json # -log
sls invoke --function iface --path ./data/iface.json # -log
# -log
