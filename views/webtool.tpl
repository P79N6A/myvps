<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title> WebTools </title>
    <!-- <link href="static/img/favicon.png" rel="icon" type="image/png"> -->
    <link href="static/css/bootstrap.min.css" rel="stylesheet" type="text/css"/>
    <style>
      .row {
        margin-top: 20px;
        margin-bottom: 10px;
      }
      .container {
        margin-top: 20px;
      }
    </style>
  </head>
  <body>
    <div class="container">
        <dev class="row">
            <dev class="col">
                <input type="button" class="btn btn-primary" onclick="window.open('http://{{.hostname}}:1983/ssh/host/127.0.0.1')" value="SSH">
            </dev>
            <!-- <dev class="col">
                <input type="button" class="btn btn-primary" onclick="window.open('http://{{.hostname}}:1980')" value="Deluge">
            </dev> -->
        </dev>
        <dev class="row">
            <dev class="col">
                <input type="button" class="btn btn-primary" onclick="window.open('http://{{.hostname}}:1981')" value="Kod">
            </dev>
            <!-- <dev class="col">
                <input type="button" class="btn btn-primary" onclick="window.open('http://{{.hostname}}:1982')" value="Mldonkey">
            </dev> -->
        </dev>
        <dev class="row">
            <dev class="col">
                <input type="button" class="btn btn-primary" onclick="window.open('http://{{.hostname}}:1980')" value="Deluge">
            </dev>
        </dev>
        <dev class="row">
            <dev class="col">
                <input type="button" class="btn btn-primary" onclick="window.open('http://{{.hostname}}:1982')" value="Mldonkey">
            </dev>
        </dev>
    </div>
    <div class="container">
      <div id="status" style="color: red;"></div>
      <div id="terminal"></div>
    </div>
  </body>
</html>
