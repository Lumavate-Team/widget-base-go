var Api = function Api(){
  var service = {
    getAsync: getAsync,
    postAsync: postAsync,
    getCookie: getCookie
  };
  return service;
  function getAsync(url, callback){
    var token = getCookie("pwa_jwt");
    var xhr = new XMLHttpRequest();
    xhr.open("GET",url, true);
    xhr.setRequestHeader("Authorization", "Bearer " + token);
    xhr.overrideMimeType("application/json");
    xhr.onreadystatechange = function() { 
        if (xhr.readyState == 4){
            callback(xhr.status, JSON.parse(xhr.responseText));
        }
    }
    xhr.send(null);
  }
  function postAsync(url, formData, callback){
    var token = getCookie("pwa_jwt");
    var xhr = new XMLHttpRequest();
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Authorization", "Bearer " + token);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.onreadystatechange = function() { 
        if (xhr.readyState == 4){
            callback(xhr.status, xhr.responseText);
        }
    }
    xhr.send(formData);
  }
  function getCookie(name){
    var cookies = document.cookie.split(";");
    for(var i=0,len=cookies.length; i < len; i++){
      var cookie = cookies[i].split("=");
      if(cookie[0].trim() == name){
        return cookie[1].trim();
      }
    }
  }
}

module.exports = Api;