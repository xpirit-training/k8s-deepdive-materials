const apiUrl = backendUrl + '/documents';

// The rest of your code remains unchanged...

async function createDocument() {
    const data = document.getElementById('docData').value;
    const response = await fetch(apiUrl, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ data })
    });
    const result = await response.json();
    console.log(result);
}



async function readDocument() {
    const id = document.getElementById('docId').value;
    const response = await fetch(`${apiUrl}/${id}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    });
    const result = await response.json();
    console.log(result);
}

async function updateDocument() {
    const id = document.getElementById('docId').value;
    const data = document.getElementById('docData').value;
    const response = await fetch(`${apiUrl}/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ data })
    });
    const result = await response.json();
    console.log(result);
}

async function deleteDocument() {
    const id = document.getElementById('docId').value;
    const response = await fetch(`${apiUrl}/${id}`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json'
        }
    });
    const result = await response.json();
    console.log(result);
}