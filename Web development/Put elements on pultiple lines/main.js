const monsterFactory = (name, age, energySource, catchPhrase) => {
    return ({name: name, age: age, energySource: energySource, catchPhrase: catchPhrase, scare() {
            console.log(catchPhrase);
        }});
}

const ghost = monsterFactory("Luisa", 251, "ectoplasm", "Boo")
ghost.scare()