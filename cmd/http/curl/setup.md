# Setup Curl

```bash
# Create AppConfig
curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X POST \
-H 'Authorization: Bearer ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b' \
-d '{"document":{"id":"b69aa108-f12e-4fa0-bf4f-ba002c11a670","name":"Woong"}}' \
https://localhost:49001/v1/woong/appconfig

# Create Home
curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X POST \
-H 'Authorization: Bearer ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b' \
-d '{"document":{"id":"home0001-f12e-4fa0-bf4f-ba002c11a670","app_config_id":"b69aa108-f12e-4fa0-bf4f-ba002c11a670","name":"Home","short_notice_list":[{"id":"shortno1-f12e-4fa0-bf4f-ba002c11a671","home_id":"home0001-f12e-4fa0-bf4f-ba002c11a670","img_url":"assets/images/icons/2x/outline_notification_important_black_24dp.png","name":"Delivery","description":"About delivery"}],"main_promotion_list":[{"id":"mainpro1-f12e-4fa0-bf4f-ba002c11a670","home_id":"home0001-f12e-4fa0-bf4f-ba002c11a670","img_url":"https://images.unsplash.com/photo-1522205408450-add114ad53fe?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=368f45b0888aeb0b7b08e3a1084d3ede&auto=format&fit=crop&w=1950&q=80","name":"Main Promotion 1","description":"Main Promotion 1","tags":[{"id":"tags0001-f12e-4fa0-bf4f-ba002c11a670","name":"new"}]},{"id":"mainpro2-f12e-4fa0-bf4f-ba002c11a670","home_id":"home0001-f12e-4fa0-bf4f-ba002c11a670","img_url":"https://images.unsplash.com/photo-1522205408450-add114ad53fe?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=368f45b0888aeb0b7b08e3a1084d3ede&auto=format&fit=crop&w=1950&q=80","name":"Main Promotion 2","description":"Main Promotion 2","tags":[{"id":"tags0001-f12e-4fa0-bf4f-ba002c11a670","name":"new"}]}]}}' \
https://localhost:49001/v1/woong/home


# Create Home Group Products
curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X POST \
-H 'Authorization: Bearer ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b' \
-d '{"documents":[{"id":"homegrp1-f12e-4fa0-bf4f-ba002c11a670","home_id":"home0001-f12e-4fa0-bf4f-ba002c11a670","group_id":"group001-f12e-4fa0-bf4f-ba002c11a670"},
{"id":"homegrp2-f12e-4fa0-bf4f-ba002c11a670","home_id":"home0001-f12e-4fa0-bf4f-ba002c11a670","group_id":"group002-f12e-4fa0-bf4f-ba002c11a670"}]}' \
https://localhost:49001/v1/woong/home/mainproducts
```

```json
// app config
{"id":"b69aa108-f12e-4fa0-bf4f-ba002c11a670","name":"Woong"}

// home
{"id":"home0001-f12e-4fa0-bf4f-ba002c11a670","app_config_id":"b69aa108-f12e-4fa0-bf4f-ba002c11a670","name":"Home"}

// short notice
{"id":"shortno1-f12e-4fa0-bf4f-ba002c11a671","home_id":"home0001-f12e-4fa0-bf4f-ba002c11a670","img_url":"assets/images/icons/2x/outline_notification_important_black_24dp.png","name":"Delivery","description":"About delivery"}

// main promotion
{"id":"mainpro1-f12e-4fa0-bf4f-ba002c11a670","home_id":"home0001-f12e-4fa0-bf4f-ba002c11a670","img_url":"https://images.unsplash.com/photo-1522205408450-add114ad53fe?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=368f45b0888aeb0b7b08e3a1084d3ede&auto=format&fit=crop&w=1950&q=80","name":"Main Promotion 1","description":"Main Promotion 1","tags":[{"id":"tags0001-f12e-4fa0-bf4f-ba002c11a670","name":"new"}]}
{"id":"mainpro2-f12e-4fa0-bf4f-ba002c11a670","home_id":"home0001-f12e-4fa0-bf4f-ba002c11a670","img_url":"https://images.unsplash.com/photo-1522205408450-add114ad53fe?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=368f45b0888aeb0b7b08e3a1084d3ede&auto=format&fit=crop&w=1950&q=80","name":"Main Promotion 2","description":"Main Promotion 2","tags":[{"id":"tags0001-f12e-4fa0-bf4f-ba002c11a670","name":"new"}]}

// home with 
{"id":"home0001-f12e-4fa0-bf4f-ba002c11a670","app_config_id":"b69aa108-f12e-4fa0-bf4f-ba002c11a670","name":"Home","short_notice_list":[{"id":"shortno1-f12e-4fa0-bf4f-ba002c11a671","home_id":"home0001-f12e-4fa0-bf4f-ba002c11a670","img_url":"assets/images/icons/2x/outline_notification_important_black_24dp.png","name":"Delivery","description":"About delivery"}],"main_promotion_list":[{"id":"mainpro1-f12e-4fa0-bf4f-ba002c11a670","home_id":"home0001-f12e-4fa0-bf4f-ba002c11a670","img_url":"https://images.unsplash.com/photo-1522205408450-add114ad53fe?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=368f45b0888aeb0b7b08e3a1084d3ede&auto=format&fit=crop&w=1950&q=80","name":"Main Promotion 1","description":"Main Promotion 1","tags":[{"id":"tags0001-f12e-4fa0-bf4f-ba002c11a670","name":"new"}]},{"id":"mainpro2-f12e-4fa0-bf4f-ba002c11a670","home_id":"home0001-f12e-4fa0-bf4f-ba002c11a670","img_url":"https://images.unsplash.com/photo-1522205408450-add114ad53fe?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=368f45b0888aeb0b7b08e3a1084d3ede&auto=format&fit=crop&w=1950&q=80","name":"Main Promotion 2","description":"Main Promotion 2","tags":[{"id":"tags0001-f12e-4fa0-bf4f-ba002c11a670","name":"new"}]}]}


// main products
{"id":"homegrp1-f12e-4fa0-bf4f-ba002c11a670","home_id":"home0001-f12e-4fa0-bf4f-ba002c11a670","group_id":"group001-f12e-4fa0-bf4f-ba002c11a670"}
{"id":"homegrp2-f12e-4fa0-bf4f-ba002c11a670","home_id":"home0001-f12e-4fa0-bf4f-ba002c11a670","group_id":"group002-f12e-4fa0-bf4f-ba002c11a670"}
```


```bash
curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X GET \
-H 'Authorization: Bearer ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b' \
http://localhost:49001/v1/woong/appconfig/b69aa108-f12e-4fa0-bf4f-ba002c11a670

curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X GET \
-H 'Authorization: Bearer ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b' \
http://localhost:49001/v1/woong/home/appconfig/b69aa108-f12e-4fa0-bf4f-ba002c11a670
```