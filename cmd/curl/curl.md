
```base
curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X POST \
-H 'Authorization: Bearer ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b' \
-d '{"status":200,"document":{"login_id":"wonkwonkwonk","login_type":"id","password":{"value":"asdfasdfasdf"},"personal":{"first_name":"wonk","last_name":"sun","birth_year":2002,"birth_month":1,"birth_day":2,"gender":"M","nationality":"KOR"},"emails":[{"email":"wonk@wonk.orgg","priority":0}]}}' \
http://localhost:49001/v1/woong/appconfig

curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X GET \
-H 'Authorization: Bearer ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b' \
http://localhost:49001/v1/woong/appconfig/b69aa108-f12e-4fa0-bf4f-ba002c11a670

curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X GET \
-H 'Authorization: Bearer ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b' \
http://localhost:49001/v1/woong/home/appconfig/b69aa108-f12e-4fa0-bf4f-ba002c11a670
```