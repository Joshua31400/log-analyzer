const dropArea = document.getElementById('drop-area');
const fileInput = document.getElementById('fileInput');

['dragenter', 'dragover', 'dragleave', 'drop'].forEach(eventName => {
    dropArea.addEventListener(eventName, preventDefaults, false);
});

function preventDefaults(e) {
    /*
    This function is used to prevent the browser from opening the file when it is dropped
     */
    e.preventDefault();
    e.stopPropagation();
}

dropArea.addEventListener('dragenter', highlight, false);
dropArea.addEventListener('dragover', highlight, false);
dropArea.addEventListener('dragleave', unhighlight, false);
dropArea.addEventListener('drop', unhighlight, false);

function highlight() {
    dropArea.style.backgroundColor = '#e9ecef';
    dropArea.style.borderColor = '#86b7fe';
}

function unhighlight() {
    dropArea.style.backgroundColor = 'transparent';
    dropArea.style.borderColor = '#ffb6ff';
}

dropArea.addEventListener('drop', handleDrop, false);
dropArea.addEventListener('click', () => fileInput.click());
fileInput.addEventListener('change', handleChange);

function handleDrop(e) {
    const dt = e.dataTransfer;
    const files = dt.files;
    handleFiles(files);
}

function handleChange(e) {
    const files = e.target.files;
    handleFiles(files);
}

function handleFiles(files) {
    /*
    This function is used to handle the file(s) that were dropped or selected
     */
    const formData = new FormData();
    formData.append('logfile', files[0]);

    fetch('/upload', {
        method: 'POST',
        body: formData
    })
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                window.location.href = '/?file=' + files[0].name;
            } else {
                alert('Upload failed');
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Upload failed');
        });
}
