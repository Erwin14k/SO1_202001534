<html>
  <head>
    <meta charset="UTF-8">
  </head>
  <body>
    <h2>**Erwin Fernando Vásquez Peñate</h2>
    <h2>**202001534</h2>
    <h2>**Sistemas Operativos 01 Sección A</h2>
    <h1>Práctica 01</h1>
    <h2>Arquitectura</h2>
    <p>
      La aplicación constará de la siguiente arquitectura:
    </p>
    <div style="text-align: center;">
      <img src="./Images/architecture.png" alt="Backbone Network">
    </div>
    <h2>FrontEnd</h2>
    <p>
      El frontend está desarrollado en react, la aplicación consta de una calculadora
      la cual puede usarse para resolver distintas operaciones matemáticas, todas
      las operaciones realizadas aquí serán enviadas al backend para ser operadas y 
      así poder obtener el resultado esperado.
    </p>
    <h2>Calculadora</h2>
    <p>
      La calculadora permite operaciones con números enteros, además de soportar
      operaciones con números negativos, las operaciones matemáticas disponibles son:
    </p>
    <ul>
      <li>División</li>
      <li>Multiplicación</li>
      <li>Suma</li>
      <li>Resta</li>
    </ul>
    <div style="text-align: center;">
      <img src="./Images/calc.png" alt="Backbone Network" >
    </div>
    <h2>Tabla De Logs</h2>
    <p>
      Esta tabla mostrará de manera detallada todas las operaciones a lo largo del
      tiempo, para llevar un historial de logs.
    </p>
    <div style="text-align: center;">
      <img src="./Images/table.png" alt="Backbone Network" >
    </div>
    <h2>Backend</h2>
    <p>
      El backend está desarrollado en Golang, el cual se encarga de hacer la conexión
      a la base de datos para todas las consultas necesarias, además de levantar un
      servidor para atender las peticiones del frontend.
    </p>
    <h2>Base De Datos</h2>
    <p>
      La aplicación necesita poder llevar un historial de los logs a lo largo del tiempo
      por lo que es necesario una base de datos para poder llevar ese control de la 
      información, MySQL es la base de datos implementada para esta solución.
    </p>
    <h2>Script SQL</h2>
    <div style="text-align: center;">
      <img src="./Images/sqlscript.png" alt="Backbone Network">
    </div>
    <h2>Docker</h2>
    <p>
      La parte de docker es fundamental a la hora de poder construir una aplicación, ya
      que mediante los contenedores e imágenes que nos ofrece docker podemos utilizar 
      aplicaciones sin necesidad de instalarlas como tal, ahorrándonos mucho tiempo
      y también recursos en nuestro ordenador, para esta aplicación se generaron 3 
      imágenes, una para el frontend, otra para el backend y la última para ejecutar el script, los enlaces de las 
      mismas se encuentran a continuación:
    </p>
    <ul>
      <li><a href="https://hub.docker.com/repository/docker/erwin14k/backend_practica1_202001534/general">Backend Calculator in Docker Hub</a></li>
      <li><a href="https://hub.docker.com/repository/docker/erwin14k/frontend_practica1_202001534/general">Frontend Calculator in Docker Hub</a></li>
      <li><a href="https://hub.docker.com/repository/docker/erwin14k/script_practica1_202001534/general">Reports Script in Docker Hub</a></li>
    </ul>
    <p>
      Además se utilizó una imágen para la base de datos, se utilizó la oficial
      de MySQL en su versión 8.0.32, la documentación oficial se encuentrá a continuación:
    </p>
    <ul>
      <li><a href="https://hub.docker.com/_/mysql">MySQL in Docker Hub</a></li>
    </ul>
    <p>
      Para la base de datos se creó adicionalmente un volumen, el cuál almacenará
      toda la información, incluso si el contenedor está detenido.
    </p>
    <p>
      Además es posible ingresar a interactuar en los
      contenedores mediante una terminal, aquí te muestro
      un ejemplo de como ingresar a la base de datos de esta aplicación en concreto:
    </p>
    <ul>
      <li>Acceder al contenedor de la base de datos: 
        <ul>
          <li>docker exec -it practice_01_db_1 mysql -uroot -pcalcpass</li>
        </ul>
      </li>
    </ul>
    <h2>Docker-Compose</h2>
    <p>
      Docker-Compose nos permite de manera rápida y eficiente poder darle vida
      a múltiples contenedores con un solo script, de manera que no haría falta
      irlos levantando uno por uno, el archivo para ejecutar el docker-compose, 
      debe ser con extensión ".yaml", los comandos tanto para iniciar un
      docker-compose como para terminarlo se muestran a continuación:
    </p>
    <ul>
      <li>Levantar Contenedores: docker-compose up -d</li>
      <div style="text-align: center;">
        <img src="./Images/compose1.png" alt="Backbone Network">
      </div>
      <div style="text-align: center;">
        <img src="./Images/compose2.png" alt="Backbone Network">
      </div>
      <li>Detener Contenedores: docker-compose down</li>
      <div style="text-align: center;">
        <img src="./Images/down.png" alt="Backbone Network">
      </div>
    </ul>
    <h2>Bash Container</h2>
    <p>
      El contenedor de bash, es necesario para ejecutar un script, el cual
      funciona para llevar un control en forma de reportes, este contenedor
      comparte un volumen con el contenedor de backend, por lo cual, al haber 
      una actualización en el archivo a analizar de parte del backend, este 
      contenedor también puede ver reflejados los cambios y mostrar reportes
      de manera actualizada, algunos comandos que fueron utilizados para dejar
      configurado de manera correcta este contenedor son:
    </p>
    <ul>
      <li>Copiar el script al volumen del contenedor: 
        <ul>
          <li>docker cp reportsScript.sh practice_01_linux_container_1:/app/Backend</li>
        </ul>
      </li>
      <li>Acceder al contenedor bash: 
        <ul>
          <li>docker exec -it practice_01_linux_container_1 /bin/sh</li>
        </ul>
      </li>
      <li>Darle permisos de ejecución al script: 
        <ul>
          <li>chmod +x reportsScript.sh</li>
        </ul>
      </li>
      <li>Ejecutar el script: 
        <ul>
          <li>./reportsScript.sh</li>
        </ul>
      </li>
    </ul>
  </body>
</html>
