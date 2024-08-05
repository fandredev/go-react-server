# Run Postgres/pgadmin with Docker
Note: please, check the .env.example file to change variables
``` 
docker-compose up
```

# Screenshots:

![image](https://github.com/user-attachments/assets/9399692d-17c1-4425-a553-ba590868924c)

### Note: 
- the pgadmin can slowly initialize. wait for between 30 and 50 seconds to start.
- Open http://localhost:8081/ to navigate to postgres page. Use your login defined in .env (check the .env.example before)
![image](https://github.com/user-attachments/assets/e40a5305-54c6-43c8-b4f3-1f695489dabf)
<br>

After you logged, click in Servers > Register > Server:

![image](https://github.com/user-attachments/assets/d191cece-f51d-466b-85f3-6ef3f3c3e7c8)

Put a name from server:

![image](https://github.com/user-attachments/assets/c0180dc5-d907-4538-a127-f2a15285ca5e)

click in connection and view the .env.example to be able fill the fields below.:

![image](https://github.com/user-attachments/assets/ab421b83-0f4f-4e07-89a4-84c1eaca46dd)

save the connection and enjoy! 😃


# Generate queries

``` 
go generate ./...
``` 