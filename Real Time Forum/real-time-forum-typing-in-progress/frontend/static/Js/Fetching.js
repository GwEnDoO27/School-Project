//Get old messafes
export async function LoadOldMessage(from, to) {
    try {
        let data = {
            from: from,
            to: to
        };
        let response = await fetch('http://localhost:8080/load-message', {
            method: 'POST',
            body: JSON.stringify(data),
            headers: {
                'Content-Type': 'application/json'
            }
        });
        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }
        let messages = await response.json();
        //console.log(messages);
        return messages;
        // Process the messages as needed, e.g., display them in the UI
    } catch (error) {
        console.error('Error loading messages:', error);
        return [];
    }
}
//get all users
export async function fetchAllUsers() {
    try {
        const response = await fetch('http://localhost:8080/all-users', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });
        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }
        const users = await response.json();
        return users;
    } catch (error) {
        console.error('Error fetching all users:', error);
        return [];
    }
}

//get the categories
export async function fetchCat() {
    try {
        const response = await fetch('http://localhost:8080/categories', {
            method: 'GET',
        });
        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }
        const categories = await response.json();
        return categories;
    }
    catch (error) {
        console.error(error)
    }
}