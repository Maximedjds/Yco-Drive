function toggleMenu() {
    var menu = document.getElementById('menuContent');
    menu.style.display = menu.style.display === 'block' ? 'none' : 'block';
}

function showSection(sectionId) {
    document.querySelectorAll('.section').forEach(section => {
        section.style.display = 'none';
    });
    document.getElementById(sectionId).style.display = 'block';
}
