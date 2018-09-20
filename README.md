# devcache

Often when I am iterating some experimental code I end up hitting the same external API calls repeatedly.

This is a simple service that performs the API request and caches the content locally to save the constant round trips. 

Run it, and then append `http://localhost:4321?url=` to any API call. i.e.

    curl http://localhost:4321?url=https://jsonplaceholder.typicode.com/users

