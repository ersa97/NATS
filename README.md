
1. docker-compose up

2. publisher
    hit API http://localhost:8000/ to test if message worked

3. services.PublishMessage in publisher/controllers/controller.go use library from https://github.com/ersa97/nats-libs    

4. consumer
    services.SubscribeNats in main.go taken from library to connect to nats
    by command (route) it will redirect the data to the suitable switch case condition 
