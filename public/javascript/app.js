// Add click event to fact to show answer
(() => {
    const answerWrapper = document.querySelectorAll('.answer-wrapper');
    const toggleBtns = document.querySelectorAll('.answer-toggle')

    for (const ans of answerWrapper) {
        ans.style.display = 'none';
    }

    for (const btn of toggleBtns) {
        btn.addEventListener('click', (e) => {
            const answer = e.target.parentElement.nextElementSibling;
            answer.style.display = answer.style.display === 'none' ? 'block' : 'none';
        });
    }

    const editForm = document.querySelector("#form-update-fact")
    const factToEdit = editForm && editForm.dataset.factid

    console.log(editForm)
    console.log(factToEdit)

    editForm && editForm.addEventListener('submit', (event) => {
        event.preventDefault();

        const formData = Object.fromEntries(new FormData(editForm));
        console.log(formData)

        return fetch(`/fact/${factToEdit}`, {
            method: 'PATCH',
            headers: {
                'Conent-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        }).then(() => document.location.href = `/fact/${factToEdit}`)
    })
})()