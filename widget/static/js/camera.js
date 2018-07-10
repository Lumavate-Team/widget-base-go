function submitCameraData(cameraData, callback){
      
      var url = window.location.href;
      Api().postAsync(url, JSON.stringify(cameraData),function(status, response){
        if (status != 200){
          console.error(response);
        }
        callback([]);
      });
    };