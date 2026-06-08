import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 20 },  // Ramp up to 20 users
    { duration: '1m', target: 20 },   // Stay at 20 users
    { duration: '30s', target: 0 },   // Ramp down to 0
  ],
  thresholds: {
    http_req_duration: ['p(95)<500'], // 95% of requests must complete within 500ms
    http_req_failed: ['rate<0.01'],   // Less than 1% of requests fail
  },
};

export default function () {
  // Test GET endpoint
  const getRes = http.get('http://localhost:8080/api/users');
  check(getRes, {
    'GET status is 200': (r) => r.status === 200,
    'GET response time < 200ms': (r) => r.timings.duration < 200,
  });

  // Test POST endpoint
  const payload = JSON.stringify({
    name: `Test User ${__VU}`,
    email: `test${__VU}@example.com`,
    age: 25,
  });

  const postRes = http.post('http://localhost:8080/api/users', payload, {
    headers: { 'Content-Type': 'application/json' },
  });

  check(postRes, {
    'POST status is 201': (r) => r.status === 201,
    'POST response time < 300ms': (r) => r.timings.duration < 300,
  });

  sleep(1);
}