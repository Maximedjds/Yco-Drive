function toggleMenu() {
    var menu = document.getElementById('menuContent');
    menu.style.display = menu.style.display === 'block' ? 'none' : 'block';
}

function showSection(sectionId) {
    document.querySelectorAll('.section').forEach(section => {
        section.style.display = 'none';
    });
    document.getElementById(sectionId).style.display = 'block';

    if (sectionId === 'messages') fetchMessages();
    if (sectionId === 'historique') fetchHistorique();
    if (sectionId === 'compte') fetchUser();
}

function fetchMessages() {
    fetch('/api/messages')
        .then(response => response.json())
        .then(data => {
            let content = data.length 
                ? data.map(m => `<div><img src="${m.avatar}" width="40"> <strong>${m.name}:</strong> ${m.message}</div>`).join("")
                : "Vous n'avez entamé encore aucune conversation.";
            document.getElementById('messages').innerHTML = `<h2>Messages</h2>${content}`;
        });
}

function fetchHistorique() {
    fetch('/api/historique')
        .then(response => response.json())
        .then(data => {
            let list = data.map(h => `<li>${h.name} - ${h.distance}</li>`).join("");
            document.getElementById('historique').innerHTML = `<h2>Historique</h2><ul>${list}</ul>`;
        });
}

function fetchUser() {
    fetch('/api/compte')
        .then(response => response.json())
        .then(user => {
            document.getElementById('compte').innerHTML = `
                <h2>Compte</h2>
                <div class="profile">
                    <img src="${user.avatar}" alt="Photo de profil" width="100">
                    <p>Nom: ${user.nom}</p>
                    <p>Prénom: ${user.prenom}</p>
                    <p>Date de naissance: ${user.date_naissance}</p>
                    <p>Téléphone: ${user.telephone}</p>
                    <p>Adresse: ${user.adresse}</p>
                </div>`;
        });
}
