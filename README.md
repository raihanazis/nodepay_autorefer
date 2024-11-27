### ⚙️&nbsp; Instaction
<img align="center" alt="line" src="https://github.com/DalpatRathore/dalpatrathore/blob/main/assets/images/line-2.svg"> 

```
pkg update && pkg upgrade -y
```
```
pkg install python
```
```
pkg install git
```
```
pkg install nano
```
```
git clone --depth=1 https://github.com/rahatmals1/nodepay_autorefer
```
```
cd nodepay_autorefer
```
```
pip install -r requirements.txt
```

⚠️ Open Via Browser And Run Captcha Solver for get captcha token
```
python main.py
```

✨ Open New Tab & Run The app
```
python script.py
```

 💥 Full Tutorial In YouTube Channel : <a href="https://m.youtube.com/@nbprg">Click Here</a>

### ⚙️&nbsp; Captcha Solver JavaScript / use via browser

```javascript

// ==UserScript==
// @name         Captcha Token Fetcher
// @namespace    https://viayoo.com/
// @version      1.0
// @description  Fetch Captcha token, reload page, wait for 5 seconds, and run for 10000 times
// @author       rahatmals
// @match        https://app.nodepay.ai/login
// @grant        GM_xmlhttpRequest
// ==/UserScript==

(function() {
    'use strict';
    let counter = 0;
    const maxIterations = 10000;  // Set maximum iterations to 10000
    let retryCounter = 0;
    const maxRetries = 8;  // Set maximum retries to 5
    // Function to simulate click on the CAPTCHA checkbox
    function clickCaptchaCheckbox() {
        const captchaCheckbox = document.querySelector('input[name="cf-turnstile-response"]');
        if (captchaCheckbox) {
            console.log('Clicking CAPTCHA checkbox...');
            captchaCheckbox.click();
        } else {
            console.log('CAPTCHA checkbox not found.');
        }
    }
    // Function to check for CAPTCHA token and make request
    function checkCaptchaToken() {
        const captchaInput = document.querySelector('input[name="cf-turnstile-response"]');
        if (captchaInput && captchaInput.value) {
            const captchaValue = captchaInput.value;
            console.log('CAPTCHA token fetched:', captchaValue);
            const url = `http://localhost:5000/post?token=${captchaValue}`;
            GM_xmlhttpRequest({
                method: "GET",
                url: url,
                onload: function(response) {
                    console.log('Response from server:', response.responseText);
                },
                onerror: function(error) {
                    console.error('Error fetching server response:', error);
                }
            });
            // Increment the counter
            counter++;
            retryCounter = 0;  // Reset retry counter after a successful fetch
            // Check if the process should stop after 10000 iterations
            if (counter < maxIterations) {
                console.log(`Iteration ${counter} completed.`);
                // Wait for 10 seconds, then reload the page
                setTimeout(() => {
                    location.reload();  // Reload the page
                }, 2000); // 10-second delay before reload
            } else {
                console.log('Completed 10000 iterations, stopping.');
            }
        } else {
            retryCounter++;
            if (retryCounter < maxRetries) {
                console.log(`CAPTCHA token not yet available, retrying... (Attempt ${retryCounter} of ${maxRetries})`);
                clickCaptchaCheckbox();  // Click the CAPTCHA checkbox if available
                setTimeout(checkCaptchaToken, 1000);  // Retry after 1 second if CAPTCHA is not available
            } else {
                console.log(`Failed to fetch CAPTCHA token after ${maxRetries} attempts, reloading page.`);
                location.reload();  // Reload the page if maximum retries reached
            }
        }
    }
    // Start the process by calling checkCaptchaToken
    checkCaptchaToken();
})();

```
<img align="center" alt="line" src="https://github.com/DalpatRathore/dalpatrathore/blob/main/assets/images/line-2.svg">
