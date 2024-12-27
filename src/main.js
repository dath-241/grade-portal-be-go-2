// require('dotenv').config();

// Sử dụng biến môi trường
// var YOUR_REDIRECT_URI = process.env.API_URL;
// var YOUR_CLIENT_ID = process.env.YOUR_CLIENT_ID;
YOUR_CLIENT_ID = '836560995551-5joc9tjdcta8fvjg05tasigq09ojvpj2.apps.googleusercontent.com';
YOUR_REDIRECT_URI = 'http://127.0.0.1:5500/src/index.html';
// console.log(YOUR_CLIENT_ID, YOUR_REDIRECT_URI)
// Phân tích chuỗi truy vấn để kiểm tra xem trang có yêu cầu từ máy chủ OAuth 2.0 không.
var fragmentString = location.hash.substring(1);
var params = {};
var regex = /([^&=]+)=([^&]*)/g, m;
while (m = regex.exec(fragmentString)) {
  params[decodeURIComponent(m[1])] = decodeURIComponent(m[2]);
}
if (Object.keys(params).length > 0 && params['state']) {
  if (params['state'] == localStorage.getItem('state')) {
    localStorage.setItem('oauth2-test-params', JSON.stringify(params));
    sendRequest();
  } else {
    console.log('State mismatch. Possible CSRF attack');
  }
}

// Hàm để tạo giá trị trạng thái ngẫu nhiên
function generateCryptoRandomState() {
  const randomValues = new Uint32Array(2);
  window.crypto.getRandomValues(randomValues);
  const utf8Encoder = new TextEncoder();
  const utf8Array = utf8Encoder.encode(String.fromCharCode.apply(null, randomValues));
  return btoa(String.fromCharCode.apply(null, utf8Array))
    .replace(/\+/g, '-')
    .replace(/\//g, '_')
    .replace(/=+$/, '');
}

// Hàm để tạo giá trị nonce ngẫu nhiên
function generateNonce() {
  const array = new Uint32Array(5);
  window.crypto.getRandomValues(array);
  return array.join('');
}

// Nếu có access token, thử gửi yêu cầu API.
// Nếu không, bắt đầu quy trình OAuth 2.0.
async function sendRequest() {
  var params = JSON.parse(localStorage.getItem('oauth2-test-params'));
  if (params && params['access_token']) {
    try {
      const response = await fetch(`https://www.googleapis.com/oauth2/v3/userinfo?access_token=${params['access_token']}`);
      if (response.ok) {
        const userInfo = await response.json(); // Lấy thông tin người dùng
        document.getElementById('userInfo').textContent = JSON.stringify(userInfo, null, 2); // Hiển thị thông tin

        // Lấy ID Token từ params
        var idToken = params['id_token'];
        if (idToken) {
          document.getElementById('idToken').value = idToken; // Gán ID Token vào input
          document.getElementById('idTokenContainer').style.display = 'block'; // Hiện container
          await sendIdTokenToBackend(idToken); // Gửi ID Token tới backend
        }
      } else if (response.status === 401) {
        oauth2SignIn(); // Token không hợp lệ, yêu cầu người dùng cấp quyền
      }
    } catch (error) {
      console.error('Error fetching user info:', error);
    }
  } else {
    oauth2SignIn(); // Nếu không có token, bắt đầu quy trình đăng nhập
  }
}
function copyIdToken() {
  // Lấy giá trị của ID Token từ input
  const idTokenInput = document.getElementById('idToken');
  if (idTokenInput && idTokenInput.value) {
    // Sao chép giá trị vào clipboard
    navigator.clipboard.writeText(idTokenInput.value)
      .then(() => {
        alert('ID Token đã được sao chép vào clipboard!');
      })
      .catch(err => {
        console.error('Không thể sao chép token:', err);
      });
  } else {
    alert('Không có ID Token để sao chép!');
  }
}
// Gửi ID Token đến backend
async function sendIdTokenToBackend(idToken) {
  try {
    const response = await fetch('http://localhost:8080/admin/api/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
      body: JSON.stringify({ idToken }),
    });

    if (response.ok) {
      console.log('ID Token đã được gửi thành công.');
    } else {
      const errorResponse = await response.text();
      console.error('Lỗi khi gửi ID Token:', errorResponse);
    }
  } catch (error) {
    console.error('Error sending ID Token:', error);
  }
}
/*
 * Tạo biểu mẫu để yêu cầu access token từ máy chủ OAuth 2.0 của Google.
 */
function oauth2SignIn() {
  // Tạo giá trị trạng thái ngẫu nhiên và lưu vào local storage
  var state = generateCryptoRandomState();
  localStorage.setItem('state', state);

  // Tạo giá trị nonce và lưu vào local storage
  var nonce = generateNonce();
  localStorage.setItem('nonce', nonce);

  // Endpoint OAuth 2.0 của Google để yêu cầu access token
  var oauth2Endpoint = 'https://accounts.google.com/o/oauth2/v2/auth';

  // Tạo phần tử để mở endpoint OAuth 2.0 trong cửa sổ mới.
  var form = document.createElement('form');
  form.setAttribute('method', 'GET'); // Gửi như một yêu cầu GET
  form.setAttribute('action', oauth2Endpoint);

  // Các tham số để gửi đến endpoint OAuth 2.0
  var params = {
    'client_id': YOUR_CLIENT_ID,
    'redirect_uri': YOUR_REDIRECT_URI,
    'scope': 'email profile openid', // Thêm 'openid' để lấy ID Token
    'state': state,
    'nonce': nonce, // Thêm nonce
    'include_granted_scopes': 'true',
    'response_type': 'id_token token' // Yêu cầu cả ID Token và Access Token
  };
  // Thêm các tham số biểu mẫu dưới dạng giá trị đầu vào ẩn
  for (var p in params) {
    var input = document.createElement('input');
    input.setAttribute('type', 'hidden');
    input.setAttribute('name', p);
    input.setAttribute('value', params[p]);
    form.appendChild(input);
  }

  // Thêm biểu mẫu vào trang và gửi nó để mở endpoint OAuth 2.0.
  document.body.appendChild(form);
  form.submit();
}