export function highlightActiveMenuItem() {
    const tabs = document.getElementsByClassName('nav-item');

    for (let tabIndexId = 0; tabIndexId < tabs.length; tabIndexId++) {
        tabs[tabIndexId].addEventListener('click', clickTab)
    }

    function clickTab(e) {
        for (let tabId = 0; tabId < tabs.length; tabId++) {
            tabs[tabId].classList.remove('active')
        }

        e.currentTarget.classList.add('active')
    }
}
