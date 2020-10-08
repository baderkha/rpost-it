## Project : URL Shortner REST API
## Author  : Ahmad Baderkhan
## Type    : API => GO LANG 
## GITHUB  : https://github.com/baderkha/comment-me
---

### Description

- This is the api for url shortner , it has a post and a get method . 
  It was written in GO and uses mysql as a db in order to fetch the k-v for the small link to the generated link.
- The Id logic works by using uint in the database , but for the front end it's encoded in base 62 so we can get a smaller looking string. 
   - for example
      - if the DATABASE id for the link was 1000001
      - the return value to the front end client would be 4C93 , this way we can access the url as domain.com/4C93 , and as a result would be alot shorter

---

### Requirements 

- In order to run locally , you must have the following Requirements
  - docker
  - docker-compose
  - go version 1.5
---

### Getting Started
In a terminal type the following command in the project root directory : =>

- ```go install```
- ```make start```

---

### Supported API Routes


#### Creating a Domain 
- Curl

    ``` bash 
      curl --request POST \
           --url https://shrter.xyz/api/links \
           --header 'content-type: application/json' \
           --data '{"url":"<<your url here>>"}' 
    ```
 
- Sample Response

    ```json 
    {"id":"1","url":"<<your url here>>"}
    ```

#### Fetching Created Domain 
- Curl

    ```bash
      curl --request GET \
           --url https://shrter.xyz/api/links/1
    ```

- Sample Response 

    ```json 
      {"id":"1","url":"<<your url here>>"}
      ```
