import http from 'k6/http';

export default function () {
  const url = 'http://localhost:8080/listarFilmes';
  const payload = JSON.stringify({
    email: 'aaa',
    password: 'bbb',
  });

  const params = {
    timeout: '150s',
    headers: {
      'Content-Type': 'application/json',
    },
  };

  http.get(url, params);
}