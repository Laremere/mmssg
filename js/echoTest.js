function newViewModel(){
	var viewModel = this;
	this.sock = new WebSocket("ws://" + window.location.host + "/sock/");
	this.sock.onmessage = function(event){
		viewModel.result.push( event.data)
	}
	this.sock.onopen = function(event){
		viewModel.result.push("connected...")
	}
	this.sock.onerror = function(event){
		console.log(event)
	}

	this.message = ko.observable();
	this.result = ko.observableArray();
	this.send = function(){
		this.sock.send(this.message())
		this.message("")
	}
}