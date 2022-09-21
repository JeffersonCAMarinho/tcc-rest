import http from 'k6/http';

export default function () {
  const url = 'http://localhost:8080';
  const payload = JSON.stringify({
    email: 'aaa',
    password: 'bbb',
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  http.post(url, params);
}