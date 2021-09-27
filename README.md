# feature-flag
Step 1 :

```
docker run -d \
    -p 8080:8080 \
    -p 9000:9000 \
    -v $HOME/flipt:/var/opt/flipt \
    markphelps/flipt:latest
```
    
 step 2: 
```
open localhost:8080 it will open Flipt UI 
```

 step 3: 
```
create new flag
```

 step4 : 
```
 mention flag name in main.go
```
