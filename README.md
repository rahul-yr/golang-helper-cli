### golang-helper-cli
    go build -ldflags "-s -w" -o rr
    sudo ls /usr/local/bin
    sudo cp rr /usr/local/bin
    rr --create-go rahul-yr/golang-test
    rr enable-addons firebase-auth
    rr enable-addons mongodb
    rr enable-addons redis
    rr enable-addons pubsub

    
### local
    go run . enable-addons firebase-auth
