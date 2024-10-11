import http from 'k6/http';
import { check } from 'k6';

// // Stress Testing
// export const options = {
//     vus: 10,
//     //duration: '5m',
//     thresholds: {
//         http_req_failed: ['rate<0.01'],
//         http_req_duration: ['p(99)<200'],
//     },
//     stages: [
//         // Level 1
//         { duration: '1m', target: 100 },
//         { duration: '2m', target: 100 },

//         // Level 2
//         { duration: '1m', target: 200 },
//         { duration: '2m', target: 200 },

//         // Level 3
//         { duration: '1m', target: 500 },
//         { duration: '2m', target: 500 },

//         // Cool down
//         { duration: '1m', target: 0 },
//     ],
// };

// Spike Testing
// export const options = {
//     vus: 10,
//     //duration: '5m',
//     thresholds: {
//         http_req_failed: ['rate<0.01'], 
//         http_req_duration: ['p(99)<200'], 
//     },
//     stages: [
//         { duration: '30s', target: 100 },

//         // Spike
//         { duration: '1m', target: 2000 },
//         { duration: '10s', target: 2000 },
//         { duration: '1m', target: 100 },

//         // Cool down
//         { duration: '30s', target: 0 },
//     ],
// };

// Soak Testing
// export const options = {
//     vus: 10,
//     //duration: '5m',
//     thresholds: {
//         http_req_failed: ['rate<0.01'], 
//         http_req_duration: ['p(99)<200'], 
//     },
//     stages: [
//         { duration: '1m', target: 200 },

//         // Sustained load over time
//         { duration: '4h', target: 200 },

//         // Cool down
//         { duration: '1m', target: 0 },
//     ],
// };

// export const options = {
//     stages: [
//         { duration: '1m', target: 1000 },
//     ],
// };

const BASE_URL = 'http://localhost:3000';

export default function () {
    let loginRes = http.post(`${BASE_URL}/login`, JSON.stringify({
        user_name: 'username1',
        password: 'password',
    }), {
        headers: { 'Content-Type': 'application/json' }
    });

    check(loginRes, {
        'login status is 200': (r) => r.status === 200,
        'received JWT token': (r) => r.cookies['jwt-token'] !== undefined,
    });

    let jwtToken = loginRes.cookies['jwt-token'] && loginRes.cookies['jwt-token'][0] ? loginRes.cookies['jwt-token'][0].value : null;
    
    if (jwtToken) {
        let userId = 'b9b2bad4-01a4-4d1b-8391-d6a7bcc6d066';
        let res = http.get(`${BASE_URL}/QueryPersonDataByUserID?user_id=${userId}`, {
            headers: {
                'Authorization': `Bearer ${jwtToken}`,
                'Content-Type': 'application/json',
            },
        });

        //console.log(res.body);

        check(res, {
            'status is 200': (r) => r.status === 200,
            'response contains user data': (r) => r.json() && r.json().user !== undefined,
        });
    } else {
        console.error('JWT token not found');
    }
}
