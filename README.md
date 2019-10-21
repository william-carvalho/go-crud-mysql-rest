# go-crud-mysql-rest

## Usage

### Installation 
    https://golang.org/doc/install
### Run
    configure DATABASE_URL=mysql://root:new_password@localhost/posts
    execute schema/init.sql
    go run main.go
then

    curl -d '{"id":"1", "title":"Black Panther", "content":"wakanda"}' -H "Content-Type: application/json" -X POST http://localhost:8005/posts
    
# Performances with [wrk](https://github.com/wg/wrk)

    sudo apt-get install build-essential libssl-dev git -y
    git clone https://github.com/wg/wrk.git wrk
    cd wrk
    make
    # move the executable to somewhere in your PATH, ex:
    sudo cp wrk /usr/local/bin
then

    node app.js

then

    wrk -d1m http://localhost:8005/posts
