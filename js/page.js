function pageReady(){
	var sock = new WebSocket("ws://" + window.location.host + "/sock/");
	sock.onerror = function(event){
		console.log(event)
	}

	//Prevent scrolling behavior
	document.ontouchstart = function(e){ 
    	e.preventDefault(); 
	}
	document.touchmove = function(e){ 
    	e.preventDefault(); 
	}

	function processOrientation(orientData){

	    var x = Math.round(orientData.x);
	    var y = Math.round(orientData.y);

	    if (x*x + y*y > 400){
	    	var rot = Math.atan2(y,x)
	    	x = Math.cos(rot) * 20
	    	y = Math.sin(rot) * 20
	    }

	    var width = $(window).width();
	    var height = $(window).height();

	    var minDim = Math.min(width, height);
	    $("#outer").attr("cx", width / 2)
	    $("#outer").attr("cy", height / 2)
	    $("#outer").attr("r", minDim / 2)

	    $("#inner").attr("cx", width / 2 + x * minDim / 50)
	    $("#inner").attr("cy", height / 2 + y * minDim / 50)
	    $("#inner").attr("r", minDim / 10)
	    sock.send(x / 20 + " " + y / 20)
	}

	window.addEventListener("MozOrientation", function(orientData){
	    var obj = {};
	    obj.x = orientData.x * 90;
	    obj.y = orientData.y * 90;
	    processOrientation(obj);
	}, true);

	window.addEventListener("deviceorientation", function(orientData) {
	    var obj = {};
	    obj.x = orientData.gamma;
	    obj.y = orientData.beta;
	    processOrientation(obj);
	}, true);
}