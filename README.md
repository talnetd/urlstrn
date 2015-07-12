##strn 

This will become a url shortener service implemented in Go lang and backed by redis. I hope I would have a lot of fun exploring go.

### Stories
1. create a shortened version of inputed url
2. Create a unique id ( simply ran string ) inside redis 
3. Store the url value in string
4. receive a request in the pattern of /:unique_id
5. process and find the url
6. if the url is found redirect the user

### Concept
* main is entry and attach requests to accessive handlers
* handlers* control and manage data and also validate before passing to datastore
* datastore* collect and analys the data and then store in the backend
* datastore* can query the necessary data for handlers so that it can give it out to the user (platform)

### Deploy

( That's all for now to make it work )
