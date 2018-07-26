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

    <link rel="stylesheet" href="{{.CacheKey}}/static/css/add.css">
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons"
      rel="stylesheet">

  </head>

  <body>
    <header>
        <lumavate-header font-color="white" background-color="#23516A" show-back-button=true Text="Add Images to Album"></lumavate-header>
    </header>

    {{ if .visible }}
        <div class="outer">
            {{ range .images}}
                <div class="border">
                    <div class="inner">
                            <div class="allImages" x-data="{{.Public_Id}}"><img src="{{.SecureUrl}}"></div>
                    </div>
                </div>
            {{end}}
        </div>
        <div id="addImages" onclick="addPhotosToAlbum()">
            <div class="addButton">
                <p id="addPhotos">Add selected images to album?</p>
            </div>
        </div>

    {{ else }}
        <div id="none">
            <div>You must take a photo before<br>you can add to an album.</div>
        </div>
    {{ end }}

    <footer>
    </footer>

    <script type="text/javascript" src="{{.CacheKey}}/lc/lumavate-components.js"></script>
    <script
        src="https://code.jquery.com/jquery-3.3.1.slim.min.js"
        integrity="sha256-3edrmyuQ0w65f8gfBsqowzjJe2iM6n0nKciPUp8y+7E="
        crossorigin="anonymous"></script>
    <script type="text/javascript">
      window.cacheKey = {{.CacheKey}};

      var imageHighlight = 'selectImage';
      var $images = $('.allImages').on("click", function(e) {
        $(this).toggleClass(imageHighlight);
        if ($('.' + imageHighlight).length > 0) {
            document.getElementById('addImages').style.display ="flex";
        } else {
            document.getElementById('addImages').style.display ="none";
        }
      });

      function getCookie(name) {
        var re = new RegExp(name + "=([^;]+)");
        var value = re.exec(document.cookie);
        return (value != null) ? unescape(value[1]) : null;
      }

      var urlParams = new URLSearchParams(window.location.search);
      document.getElementsByTagName("lumavate-header")[0].setAttribute("Text", urlParams.get('album'));
      document.getElementById('addPhotos').innerHTML = "Add photos to " + urlParams.get('album') + " album?";

      function addPhotosToAlbum() {
        var images = []
        var list = document.querySelectorAll("div.selectImage");
        for (var i = 0; i < list.length; i++) {
          var picture = list[i].getAttribute('x-data');
          var fields = picture.split('/');
          var final = fields[1];
          images.push(final);
        }

        var payload = {};
        payload['albumTitle'] = urlParams.get('album');
        payload['images'] = images;
        console.log(JSON.stringify(payload));

        const http = new XMLHttpRequest();
        http.open("PUT", "{{ .taskURL }}/cloudinary/images/update", true);
        http.onreadystatechange = function() {
          if(http.readyState === 4 && http.status === 200) {
            console.log(http.readyState);
            console.log(http.responseText);
            setTimeout(function () {
              location.href = "{{.wid}}";
            }, 1500);
          }
        }

        http.setRequestHeader('Authorization', 'Bearer' + ' ' + getCookie("pwa_jwt"));
        http.setRequestHeader('Content-Type', 'application/json');
        http.setRequestHeader('Access-Control-Allow-Headers', 'Origin, X-Requested-With, Content-Type, Accept');
        http.setRequestHeader('Access-Control-Allow-Origin', '*');
        http.setRequestHeader('Access-Control-Allow-Methods', '*');

        http.send(JSON.stringify(payload));
      }

    </script>
  </body>
</html>