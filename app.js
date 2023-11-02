fetch('http://localhost:1211/result')
      .then(response => response.json())
      .then(json => console.log(json))