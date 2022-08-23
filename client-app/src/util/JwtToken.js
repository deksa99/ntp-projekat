export function getAccIdFromToken() {
  try {
    let token = JSON.parse(localStorage.getItem("user")).Token;
    let parsedToken = JSON.parse(atob(token.split(".")[1]));
    let id = parsedToken.id;
    return id;
  } catch (e) {
    return null;
  }
}

export function getUserIdFromToken() {
  try {
    let token = JSON.parse(localStorage.getItem("user")).Token;
    let parsedToken = JSON.parse(atob(token.split(".")[1]));
    let user_id = parsedToken["user_id"];
    return user_id;
  } catch (e) {
    return null;
  }
}

export function getRoleFromToken() {
  try {
    let token = JSON.parse(localStorage.getItem("user")).Token;
    let parsedToken = JSON.parse(atob(token.split(".")[1]));
    let role = parsedToken.role;
    return role;
  } catch (e) {
    return null;
  }
}

export function getServiceFromToken() {
  try {
    let token = JSON.parse(localStorage.getItem("user")).Token;
    let parsedToken = JSON.parse(atob(token.split(".")[1]));
    let sid = parsedToken["service_id"];
    return sid;
  } catch (e) {
    return null;
  }
}
