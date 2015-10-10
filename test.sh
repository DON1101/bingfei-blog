echo "Testing Go application..."

echo "Creating test database..."
psql -c 'create database bingfei_test;' -U postgres
export ORM_DRIVER="postgres"
export ORM_SOURCE="user=postgres dbname=bingfei_test sslmode=disable"

echo "Start unit testing..."
go test test/...

# echo "Deleting test database..."
# psql -c "drop database bingfei_test" -U postgres

echo "Finished"