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
						'content-type': 'multipart/form-data',
						'responseType': 'blob'
					}
				};
				axios
					.post('http://localhost:1323/api/fileupload?splitNumber=' + this.splitNumber, formData, config)
					.then(function(response) {
						const url = window.URL.createObjectURL(new Blob([response.data]));
						const link = document.createElement('a');
						link.href = url;
						link.setAttribute('download', 'file.tsv');
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
