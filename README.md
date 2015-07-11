~~##strn~~ 

This will become a url shortener service implemented in Go lang and backed by redis. I hope I would have a lot of fun exploring go.

### Stories
1. create a shortened version of inputed url
⋅⋅* Create a unique id ( int simply say ) insid redis 
..* Store the url value in string
2. receive a request in the pattern of /:unique_id
..* process and find the url
..* if the url is found redirect the user

( That's all for now to make it work )
