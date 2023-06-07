function base64UrlEncode(arrayBuffer) {
  if (!arrayBuffer || arrayBuffer.length === 0) {
    return undefined;
  }

  const u8a = typeof arrayBuffer === 'string' ? new TextEncoder().encode(arrayBuffer) : new Uint8Array(arrayBuffer)
  return btoa(String.fromCharCode.apply(null, u8a))
    .replace(/=/g, "")
    .replace(/\+/g, "-")
    .replace(/\//g, "_");
}

function base64UrlDecode(base64url) {
  let input = base64url
    .replace(/-/g, "+")
    .replace(/_/g, "/");
  let diff = input.length % 4;
  if (!diff) {
    while(diff) {
      input += '=';
      diff--;
    }
  }

  return Uint8Array.from(atob(input), (c) => c.charCodeAt(0));
}

function removeEmpty(obj) {
  for (let key in obj) {
    if (obj[key] == null || obj[key] === "") {
      delete obj[key];
    } else if (typeof obj[key] === 'object') {
      removeEmpty(obj[key]);
    }
  }
}
