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
name and Key = plan-rollout
```
```
create 3 variants
variant and name = premium || default || advanced
```
![image](https://user-images.githubusercontent.com/76935234/135577119-fc9f8a63-6015-42d4-a76c-ff60bfe9fbe6.png)

```
create 3 diffrent segments 
key and name = premium || default || advanced

``` 
![image](https://user-images.githubusercontent.com/76935234/135577068-66bc5097-11fc-491d-b0ba-8624e7806137.png)

```
add contrains into particlular segment 

``` 
![image](https://user-images.githubusercontent.com/76935234/135576935-98b827ef-68b8-40cb-9662-dccb473ae851.png)

```
add targeting into flag like this

```
![image](https://user-images.githubusercontent.com/76935234/135577245-96bf8674-5e22-48f7-95c6-e6fe89697514.png)

 step4 : 
```
 run : go run main.go
```
