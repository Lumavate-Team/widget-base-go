<!DOCTYPE html>

<html lang="en">
  <head>
    <title>
        {{.data.InstanceName}}
    </title>

    <meta charset="utf-8">
    <base href="{{.WidgetUrlPrefix}}">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width,initial-scale=1.0">
    <meta name="description" content="Basic Go Widget">
    <meta name="theme-color" content="#ffffff">
    <noscript>Javascript required for this site to work.</noscript>

    <link rel="apple-touch-icon" sizes="180x180" href="/iot/favicon-180x180.png">
    <link rel="icon" type="image/png" sizes="32x32" href="/iot/favicon-32x32.png">
    <link rel="icon" type="image/png" sizes="16x16" href="/iot/favicon-16x16.png">
    <link rel="manifest" href="/manifest.json">
    <link rel="shortcut icon" href="/iot/favicon.ico">

    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="apple-mobile-web-app-status-bar-style" content="black">

    <link rel="stylesheet" href="https://unpkg.com/material-components-web@latest/dist/material-components-web.min.css">
    <link rel="stylesheet" href="{{.CacheKey}}/static/css/album.css">
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons"
      rel="stylesheet">

  </head>

  <body>
    <header>
        <lumavate-header font-color="white" background-color="#23516A" show-back-button=true Text="Your Album"></lumavate-header>
    </header>

    <div id="none">
      <div>There are no pictures in<br>this album yet.</div>
    </div>

    <div id="div1"></div>

    <div class="add">
        <button onclick="addPhotos()" class="mdc-fab" aria-label="Add">
            <span class="mdc-fab__icon material-icons">add</span>
        </button>
    </div>

    <footer>

    </footer>

    <script type="text/javascript" src="{{.CacheKey}}/lc/lumavate-components.js"></script>
    <script type="text/javascript">
      window.cacheKey = {{.CacheKey}};
      document.body.onload = addElement;

      function getCookie(name) {
        var re = new RegExp(name + "=([^;]+)");
        var value = re.exec(document.cookie);
        return (value != null) ? unescape(value[1]) : null;
      }

      var urlParams = new URLSearchParams(window.location.search);

      const http = new XMLHttpRequest()
      http.open("GET", "{{ .taskURL }}/cloudinary/images?album=" + urlParams.get('album'));
      http.setRequestHeader('Authorization', 'Bearer' + ' ' + getCookie("pwa_jwt"));
      http.setRequestHeader('Access-Control-Allow-Headers', 'Origin, X-Requested-With, Content-Type, Accept');
      http.setRequestHeader('Access-Control-Allow-Origin', '*');
      http.setRequestHeader('Access-Control-Allow-Methods', 'GET');
      http.send()

      function addElement() {
          var pictures = JSON.parse(http.responseText);
          var allPictures = pictures.allImages;
          document.getElementsByTagName("lumavate-header")[0].setAttribute("Text", pictures.albums[0]);
          console.log(allPictures.length);

          if (allPictures.length === 0) {
            document.getElementById('none').style.display ="block";
          } else {
            var outer = document.createElement("div");
            outer.className = "outer";
            for (var i in allPictures) {
                var border = document.createElement("div");
                border.className = "border";
                var inner = document.createElement("div");
                inner.className = "inner";
                var allImages = document.createElement("div");
                allImages.className = "allImages";
                var newIMG = document.createElement("IMG");
                newIMG.src = allPictures[i].secure_url;
                allImages.appendChild(newIMG);
                inner.appendChild(allImages);
                border.appendChild(inner);
                outer.appendChild(border);
                // add the newly created element and its content into the DOM 
                var currentDiv = document.getElementById("div1");
                document.body.insertBefore(outer, currentDiv);
            }
          }
      }

      function addPhotos() {
        window.location = '{{ $.WidgetInstancePrefix }}/add?album=' + urlParams.get('album');
      }
    </script>
  </body>
</html>