import 'phaser';
import CursorKeys = Phaser.Types.Input.Keyboard.CursorKeys;
import Vector2 = Phaser.Math.Vector2;
import {SpaceObject} from "./spaceObject";
import {SpaceShip} from "./spaceShip";
import {Rocket} from "./rocket";
import ParticleEmitterManager = Phaser.GameObjects.Particles.ParticleEmitterManager;

const config = {
    type: Phaser.AUTO,
    width: window.innerWidth,
    height: window.innerHeight,
    scene: {
        preload: preload,
        create: create,
        update: update
    }
};

const _game = new Phaser.Game(config);
let cursors: CursorKeys;
let ship: SpaceObject;
let rocket: Rocket;
let rocket2: Rocket;
let particleEmitterManager: ParticleEmitterManager;

function preload() {
    this.load.atlas('shapes', 'assets/shapes.png', 'assets/shapes.json');
}

function create() {
    (window as any).game = this;
    particleEmitterManager = this.add.particles('shapes');

    cursors = this.input.keyboard.createCursorKeys();
    ship = new SpaceShip(this, 200, 200);
    rocket = new Rocket(this, window.innerWidth / 2, window.innerHeight / 2, particleEmitterManager);
    rocket2 = new Rocket(this, window.innerWidth / 2, window.innerHeight / 2, particleEmitterManager, 0xff0000);
}

function update() {

    if (cursors.up?.isDown) {
        ship.moveForward();
        //particleEmillter.emitParticle(10, 200, 200);
    }

    if (cursors.left?.isDown) {
        ship.rotate(-1);
    } else if (cursors.right?.isDown) {
        ship.rotate(1);
    }

    rocket.rotate(1);
    rocket.moveForward();

    rocket2.rotate(-1);
    rocket2.moveForward();
}