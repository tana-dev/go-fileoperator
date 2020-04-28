(function(){
	new Vue({
		el: '#app',
		data: {
			uploadFile: null,
			splitNumber: null
		},
		methods: {
			selectedFile: function(e) {
				e.preventDefault();
				let files = e.target.files;
				this.uploadFile = files[0];
			},
			inputSplitNumber: function(e) {
				this.splitNumber = e.target.value;
			},
			upload: function() {
				let formData = new FormData();
				formData.append('file', this.uploadFile);
				let config = {
					headers: {
						'content-type': 'multipart/form-data'
					}
				};
				axios
					.post('http://localhost:1323/api/fileupload?splitNumber=' + this.splitNumber, formData, config)
					.then(function(response) {
						// response 処理
						console.log("sucess");
					})
					.catch(function(error) {
						console.log("error");
					})
			}
		}
	});
})();
