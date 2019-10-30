# add
curl -v -X POST -d '{"title":"abc", "isbn":"1234"}' localhost:8080/api/books
curl -v -X POST -d '{"title":"abc", "isbn":"1234567"}' localhost:8080/api/books

# list
curl -X GET localhost:8080/api/books | jq

# get by isbn
curl -X GET localhost:8080/api/book/1234 | jq
curl -X GET localhost:8080/api/book/1234567 | jq

# update by isbn
curl -v -X PUT -d '{"title":"1234 book", "isbn":"1234"}' localhost:8080/api/book/1234
# list
curl -X GET localhost:8080/api/books | jq

# delete by isbn
curl -v -X DELETE localhost:8080/api/book/1234567
# list
curl -X GET localhost:8080/api/books | jq