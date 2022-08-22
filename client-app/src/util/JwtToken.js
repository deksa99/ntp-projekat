export function getIdFromToken() {
    try {
      let token = JSON.parse(localStorage.getItem("user")).Token;
      let parsedToken = JSON.parse(atob(token.split(".")[1]));
      let id = parsedToken.id;
      console.log("ID", id)
      return id;
    } catch (e) {
      return null;
    }
  }
  
  export function getRoleFromToken() {
    try {
      let token = JSON.parse(localStorage.getItem("user")).Token;
      let parsedToken = JSON.parse(atob(token.split(".")[1]));
      let role = parsedToken.role;
      console.log("ROLE", role)
      return role;
    } catch (e) {
      return null;
    }
  }
  