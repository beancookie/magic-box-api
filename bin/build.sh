docker build -t magic_box_api:1.0.0 .

docker tag magic_box_api:1.0.0 luzhong256/magic_box_api:1.0.0

docker login 

docker push luzhong256/magic_box_api:1.0.0