This project purpose for demo on Flutter interact about using swagger both client and server.

`Prerequisites:`

Download all this prerequisities before you go to next step:

    1. Server side, used golang version go1.13.4, download from https://golang.org
    2. Client side, used flutter version v1.13.2, how to setup from https://flutter.dev/docs/get-started/install
    3. Swagger spec api 2.0.0
    4. Swagger-codegen cli v2.4.10, download cli from https://github.com/swagger-api/swagger-codegen and put to client/swagger/
    5. Go-Swagger v0.21.0, download and setup from https://github.com/go-swagger/go-swagger

`Structures:`

In this project contains 3 folders: 

    1. client, contains flutter apps
    2. schema, contains swagger specification
    3. server, contains golang apps

`How to generate client swagger:`

    1. Open terminal and go to this current folder
    2. Run command:  java -jar ./client/swagger/swagger-codegen-cli-2.4.10.jar generate -l dart -i ./schema/profile-spec.yml -o ./client/swagger -DbrowserClient=false
    3. Go to client folder, open pubspec.yaml and add line on depedencies:
        swagger:
            path: ./swagger
    4. Go to /client/swagger/, open pubspec.yaml and add line after description:
        environment:
            sdk: ">=2.1.0 <3.0.0"
 
 `How to generate server swagger:`
 
    1. Open terminal and go to this current folder
    2. Run command: swagger generate server -f ../schema/profile-spec.yml
    3. Run command: go get -u -f ./...
    4. Run command: go mod vendor

`Limitation:`

    1. Only Swagger spec 2.0.0
    2. Dart 1.20.0 or later OR Flutter 0.0.20 or later