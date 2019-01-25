const tokenKey = 'token-key';

export function saveToken(jwt) {
  sessionStorage.setItem(tokenKey, jwt);
}

export function deleteToken() {
  sessionStorage.removeItem(tokenKey)
}

export function token() {
  return sessionStorage.getItem(tokenKey)
}
