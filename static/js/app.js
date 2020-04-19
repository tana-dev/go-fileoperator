(function(){
	new Vue({
		el: '#app',
		data: {
			uploadFile: null
		},
		methods: {
			selectedFile: function(e) {
				// 選択された File の情報を保存しておく
				e.preventDefault();
				let files = e.target.files;
				this.uploadFile = files[0];
			},
			upload: function() {
				// FormData を利用して File を POST する
				let formData = new FormData();
				formData.append('file', this.uploadFile);
				let config = {
					headers: {
						'content-type': 'multipart/form-data'
					}
				};
				axios
					.post('http://localhost:1323/api/fileupload', formData, config)
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
