### Optimistic Locking

```txt
> k6 run test.js

         /\      Grafana   /‾‾/
    /\  /  \     |\  __   /  /
   /  \/    \    | |/ /  /   ‾‾\
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/

     execution: local
        script: test.js
        output: -

     scenarios: (100.00%) 1 scenario, 15000 max VUs, 10m30s max duration (incl. graceful stop):
              * flash_sale: 1 iterations for each of 15000 VUs (maxDuration: 10m0s, gracefulStop: 30s)

WARN[0002] Could not get a VU from the buffer for 400ms  executor=per-vu-iterations scenario=flash_sale
ERRO[0032] Unexpected Error (503): {"error":"system busy, please try again"}  source=console
ERRO[0033] Unexpected Error (503): {"error":"system busy, please try again"}  source=console
ERRO[0033] Unexpected Error (503): {"error":"system busy, please try again"}  source=console
ERRO[0033] Unexpected Error (503): {"error":"system busy, please try again"}  source=console
ERRO[0034] Unexpected Error (503): {"error":"system busy, please try again"}  source=console


  █ TOTAL RESULTS

    checks_total.......: 15000  414.540345/s
    checks_succeeded...: 66.66% 10000 out of 15000
    checks_failed......: 33.33% 5000 out of 15000

    ✗ status is 200
      ↳  66% — ✓ 10000 / ✗ 5000

    CUSTOM
    errors_other...................: 5      0.13818/s
    errors_sold_out................: 4995   138.041935/s

    HTTP
    http_req_duration..............: avg=23.34s min=67.29ms  med=27.51s max=35.38s p(90)=34.86s p(95)=35.09s
      { expected_response:true }...: avg=17.96s min=67.29ms  med=17.52s max=34.88s p(90)=31.36s p(95)=32.6s
    http_req_failed................: 33.33% 5000 out of 15000
    http_reqs......................: 15000  414.540345/s

    EXECUTION
    iteration_duration.............: avg=22.11s min=101.01ms med=24.59s max=32.6s  p(90)=32s    p(95)=32.19s
    iterations.....................: 15000  414.540345/s
    vus............................: 5088   min=0             max=14762
    vus_max........................: 15000  min=7270          max=15000

    NETWORK
    data_received..................: 3.8 MB 105 kB/s
    data_sent......................: 3.6 MB 99 kB/s




running (00m36.2s), 00000/15000 VUs, 15000 complete and 0 interrupted iterations
flash_sale ✓ [======================================] 15000 VUs  00m32.6s/10m0s  15000/15000 iters, 1 per VU
```

### Pessimistic Locking

```txt
> k6 run test.js

         /\      Grafana   /‾‾/
    /\  /  \     |\  __   /  /
   /  \/    \    | |/ /  /   ‾‾\
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/

     execution: local
        script: test.js
        output: -

     scenarios: (100.00%) 1 scenario, 15000 max VUs, 10m30s max duration (incl. graceful stop):
              * flash_sale: 1 iterations for each of 15000 VUs (maxDuration: 10m0s, gracefulStop: 30s)



  █ TOTAL RESULTS

    checks_total.......: 15000  476.856027/s
    checks_succeeded...: 66.66% 10000 out of 15000
    checks_failed......: 33.33% 5000 out of 15000

    ✗ status is 200
      ↳  66% — ✓ 10000 / ✗ 5000

    CUSTOM
    errors_sold_out................: 5000   158.952009/s

    HTTP
    http_req_duration..............: avg=17.79s min=71.54ms  med=18.38s max=30.77s p(90)=29.84s p(95)=30.13s
      { expected_response:true }...: avg=12.18s min=71.54ms  med=11.85s max=29.89s p(90)=22.44s p(95)=24.69s
    http_req_failed................: 33.33% 5000 out of 15000
    http_reqs......................: 15000  476.856027/s

    EXECUTION
    iteration_duration.............: avg=17.2s  min=184.67ms med=19.09s max=27.85s p(90)=26.99s p(95)=27.23s
    iterations.....................: 15000  476.856027/s
    vus............................: 5039   min=0             max=14586
    vus_max........................: 15000  min=7213          max=15000

    NETWORK
    data_received..................: 3.8 MB 121 kB/s
    data_sent......................: 3.6 MB 114 kB/s




running (00m31.5s), 00000/15000 VUs, 15000 complete and 0 interrupted iterations
flash_sale ✓ [======================================] 15000 VUs  00m27.9s/10m0s  15000/15000 iters, 1 per VU
```
