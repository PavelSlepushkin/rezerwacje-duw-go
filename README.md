# Rezerwacje DUW

#### To begin reservation

1. install `golang` of min version `1.11`.

    ```$ brew install go```

2. go to the application folder.
3. copy `user.yml.template`

   ```$ cp user.yml.template user.yml``` 

4. edit `user.yml` file providing your credentials and information required for reservation.
5. run the application

   ```$ go run main.go```

6. wait until the message appears `Reservation completed for the following city Wrocław, slot 123456 and time 2018-11-03 11:15:00. Check your email or DUW site`.
7. close the app by pressing any key.
8. check your email or DUW site.



