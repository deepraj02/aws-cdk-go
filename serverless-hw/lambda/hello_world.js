export async function handler(event) {
    return {
        statusCOde:200,
        headers: {"Content-Type": "text/plain"},
        body: JSON.stringify({message: "Hello World"}),
    }
}