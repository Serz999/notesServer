### Prerequierments
go compiler ^1.12.4
docker ^20.10.24
### Quickstart
```
git clone https://github.com/serz999/notesServer.git
cd notesServer 
cp -n .env.example .env
make STORAGE=postgres # Default
```
### Set up your storage
```
make STORAGE=postgres # Default
make STORAGE=list
make STORAGE=map
```
