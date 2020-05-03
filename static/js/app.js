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
					responseType: 'blob',
					headers: {
						'content-type': 'multipart/form-data',
						Accept: 'application/zip'
					}
				};
				axios
					.post('http://localhost:1323/api/v1/filesplit?splitNumber=' + this.splitNumber, formData, config)
					.then(function(response) {
						const url = window.URL.createObjectURL(new Blob([response.data]));
						const link = document.createElement('a');
						link.href = url;
						link.setAttribute('download', 'split.zip');
						document.body.appendChild(link);
						link.click();
					})
					.catch(function(error) {
						console.log("error");
					})
			}
		}
	});
})();
