// TODO : sys round 

function initializeGame() {
    //Initaite some vars
    const container = document.querySelector('.container');
    const cursor = document.querySelector('.cursor');
    const blood = document.querySelector('.blood');
    const scoreDisplay = document.getElementById('score');
    const startbtn = document.querySelector('.btn')
    const divbullet = document.getElementById('bullet')
    const fpsdisplay = document.getElementById('fps')
    const timerdisplay = document.getElementById('timer')
    const NewWaveText = document.getElementById('NewWaveText')
    const lifedisplay = document.getElementById('Life')

    // Hide component

    NewWaveText.style.display = 'none'
    startbtn.style.display = 'none'

    // init somes var
    const nameImg = ['adrien', 'loulou', 'manu', 'mathilde', 'sanderine', 'Meluche'];
    let ducks = [];

    let score = 0;
    let lastMoveTime = 0;
    let timerWave = 10;
    let LastGame = "Start";

    let bullet = 0
    let gamePaused = false

    let fps = 0
    let deadDuck = 0

    let time = 0
    let life = 3

    // Display the life

    updateLifeDisplay()
    //Create the class duck and give style
    class Canard {
        constructor() {
            this.element = document.createElement('div');
            this.randomImage = this.getRandomPng();
            this.setImage();
            this.element.classList.add('duck');
            this.element.style.position = 'absolute';
            this.element.style.transition = '2s ease-in-out';
            this.element.style.margin = '20px';
            let posX = (window.innerWidth / 2);
            let posY = ((container.offsetHeight / 2) + 300);
            this.element.style.left = posX + 'px';
            this.element.style.top = posY + 'px';

            // add the ducks to the body
            container.appendChild(this.element);
        }

        // give the sprite at the ducks
        setImage() {
            this.element.innerHTML = `<img src="asset/flyright.gif" width="110">`;
            this.element.firstChild.addEventListener('click', () => shoot(this.element));
        }

        getRandomPng() {
            const randomIndex = Math.floor(Math.random() * nameImg.length);
            return nameImg[randomIndex];
        }

        // Give the movement at the dicks
        moveRandomly() {
            let posX = Math.random() * (window.innerWidth - 110);
            let posY = Math.random() * (container.offsetHeight - 250);
            this.element.style.left = posX + 'px';
            this.element.style.top = posY + 'px';
        }
    }

    // Function who create some ducks randomly
    function createDucks() {
        ducks = [];
        for (let i = 0; i < Math.floor(Math.random() * 10) + 1; i++) {
            ducks.push(new Canard());
        }
        deadDuck = ducks.length
        bullet = ducks.length + 2
        updateBulletDisplay()
    }
    createDucks()

    // Allow to animate 
    function animate(time) {
        fps++

        if (gamePaused) return;
        if (time - lastMoveTime > 1000) {
            ducks.forEach(duck => duck.moveRandomly());
            lastMoveTime = time;
        }
        requestAnimationFrame(animate);

    }

    requestAnimationFrame(animate);

    // Verifying  every seconds
    setInterval(() => {
        fpsdisplay.innerHTML = "FPS: " + fps;
        fps = 0
        // Check if the game is not in pause and if life are supp to 0 to adding time
        if (!gamePaused && life != 0) {
            time++
        }
        // Stop the timerwave if the game is paused
        if (gamePaused === false) {
            timerWave--;
        }

       

        // Displaying the timer
        timerdisplay.innerHTML = "Time: " + time + 's'
        // Check for retrive life and launch the next round
        if (timerWave == 0 && life != 0) {
            life--
            updateLifeDisplay()
            nextround()
        }




    }, 1000)

    // Track the mouse
    window.addEventListener('mousemove', function (e) {
        cursor.style.left = e.pageX - 15 + 'px';
        cursor.style.top = e.pageY - 40 + 'px';
    });

    // If click display the blood
    window.addEventListener('click', function (e) {
        blood.style.left = e.pageX - 85 + 'px';
        blood.style.top = e.pageY - 60 + 'px';

        // time of the blood is displayed
        setTimeout(function () {
            blood.style.display = 'none';
        }, 1200);
        // Decrease the number of bullets
        decreaseBullet()
    });


    // Function who manage the shoot
    function shoot(duck) {
        // Check the number of bullets
        if (bullet > 0) {
            blood.style.display = 'block';
            duck.style.display = 'none';
            score += 100;
            scoreDisplay.innerHTML = "Score: " + score;
            // reduce the number of ducks for comparate
            ducks = ducks.filter(d => d !== duck);
            deadDuck--;
            // If no duck launch the nextround
            if (deadDuck === 0) {
                setTimeout(() => nextround(), 1200);
            }
        }
    }


    // Fucntion who decrease the numbers of bullets
    function decreaseBullet() {
        // define the numbers of bullets at start of round
        bullet = Math.max(0, bullet - 1);
        // display the numbers of bullets
        updateBulletDisplay();
        // check if the round is loose
        if (bullet >= 0 && deadDuck === 0) {
            LastGame = "Win";
            updateLifeDisplay()
            setTimeout(() => nextround(), 1200);
        }
        // check if the user loose a life
        else if ((bullet === 0 && life > 0)) {
            LastGame = "Lose";
            life--
            updateLifeDisplay()
            nextround()

        
        }
    }

    // Manage the display of bullets
    function updateBulletDisplay() {
        divbullet.innerHTML = "Bullets: " + bullet;
    }

    // Manage the display of lifes
    function updateLifeDisplay() {
        lifedisplay.innerHTML = "Life: " + life
    }

    // MAnage the next round 
    function nextround() {
        // erased all the actual ducks
        ducks.forEach(duck => container.removeChild(duck.element));
        ducks = [];

        // If have life relaunch a round
        if (life != 0) {
            NewWaveText.style.display = "flex"
            if (LastGame == "Win") {
                NewWaveText.style.color = "green"
            } else {
                NewWaveText.style.color = "red"
            }
            setTimeout(()=>{
                NewWaveText.style.display = "none"
                NewWaveText.style.color ="white"
            },2000)
            scoreDisplay.innerHTML = "Score: " + score;
            createDucks();
            timerWave = ducks.length + 10;
            timerWave = ducks.length + 10;
            divbullet.innerHTML = "Bullets: " + bullet;
        }


        // If no life display game over menu
        if (life === 0) {
            gameOver()
        }

    }

    // Reset the games
    function reset() {
        gamePaused = false
        // Clear existing ducks
        ducks.forEach(duck => container.removeChild(duck.element));
        ducks = [];
        // Reset game state
        LastGame = "Start"
        time = 0
        score = 0;
        life = 3
        scoreDisplay.innerHTML = "Score: " + score;
        divbullet.innerHTML = "Bullets: " + bullet;
        createDucks();
    }

    // Manage the game over menu

    function gameOver() {
        gamePaused = true

        // Stopp the time
        if (gamePaused === true) {
            time === 0
        }
        // crrate the menu

        // Stopp the time
        if (gamePaused === true) {
            time === 0
        }
        // crrate the menu
        let gameovermenu = document.createElement("div");
        gameovermenu.classList.add('gameover');
        gameovermenu.innerHTML = `
            <div class="gameovermenuContent">
                <h1>Game Over</h1>
                <button class="retryBtn">Retry</button>
            </div>
        `;
        // Add to the body
        // Add to the body
        document.body.appendChild(gameovermenu);
        // Check if the retry button is click for reset the game
        document.querySelector('.retryBtn').addEventListener('click', () => {
            gameovermenu.remove();
            reset();
            requestAnimationFrame(animate)
        });
    }



    cursor.style.display = 'block';
    blood.style.display = 'none';
    scoreDisplay.innerHTML = "Score: 0";

    // Pause menu
    document.addEventListener('keydown', function (e) {
        const keyName = e.key
        // If the p key is pressed the pause menu is called
        if (keyName === "p") {
            if (!gamePaused === true) {
                gamePaused = true
                // create the div for the menu
                let startMenu = document.createElement('div')
                startMenu.classList.add('startMenu')
                startMenu.innerHTML = `
                    <div class="menuContent">
                        <h1>Game Paused</h1>
                        <button class="resumeBtn">Resume</button>
                        <button class="restartBtn">Restart</button>
                    </div>
                `;
                // add to the body
                document.body.appendChild(startMenu)
                // check if the user resume the game
                document.querySelector('.resumeBtn').addEventListener('click', () => {
                    gamePaused = false
                    startMenu.remove();
                    requestAnimationFrame(animate)
                });
                // check if the user restart the game
                document.querySelector('.restartBtn').addEventListener('click', () => {
                    startMenu.remove();
                    reset();
                    gamePaused = false;
                    requestAnimationFrame(animate);
                });
                // allow to re press p to resume the game
            } else {
                let startMenu = document.querySelector('.startMenu');
                if (startMenu) {
                    gamePaused = true
                    startMenu.remove();
                    requestAnimationFrame(animate);
                }
                gamePaused = false;
            }
        }
    })
}

// check if the start button is click for laucn the game
document.querySelector('.btn').addEventListener('click', function () {
    initializeGame();
});
