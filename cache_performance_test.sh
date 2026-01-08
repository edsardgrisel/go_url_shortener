echo "Testing cache speedup"

TEST_URL="https://google.com"
BASE_URL="http://localhost:8090"

echo "Creating shortened url..."
SHORTENED_URL=$(curl -s -X POST ${BASE_URL}/shorten -d "${TEST_URL}")
echo -e "\n Shortened url: ${SHORTENED_URL}"

echo -e "\n Clearing cache to simulate cache miss"
redis-cli DEL "url:${SHORTENED_URL}"

echo -e "\n Cache miss request:"
time curl -s -L ${BASE_URL}/r/${SHORTENED_URL} > /dev/null 2>&1

echo -e "\n Cache hit request 1:"
time curl -s -L ${BASE_URL}/r/${SHORTENED_URL} > /dev/null 2>&1

echo -e "\n Cache hit request 2:"
time curl -s -L ${BASE_URL}/r/${SHORTENED_URL} > /dev/null 2>&1

echo -e "Deleting ${SHORTENED_URL} from mysql and redis..."
redis-cli DEL "url:${SHORTENED_URL}"
mysql -u root -e "USE url_shortener; DELETE FROM urls WHERE shortened_url = '${SHORTENED_URL}';"