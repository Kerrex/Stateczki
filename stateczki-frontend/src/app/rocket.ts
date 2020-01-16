import {SpaceObject} from "./spaceObject";
import Scene = Phaser.Scene;
import ParticleEmitterManager = Phaser.GameObjects.Particles.ParticleEmitterManager;
import BlendModes = Phaser.BlendModes;

export class Rocket extends SpaceObject {

    private emitter = this.particleEmitterManager.createEmitter(
        {
            frame: 'circle_01',
            x: this.positionX,
            y: this.positionY,
            lifespan: 250,
            speed: { min: 400, max: 600 },
            angle: {
                min: 85,
                max: 95
            },
            scale: { start: 0.2, end: 0 },
            quantity: 2,
            blendMode: BlendModes.ADD
        }
    );

    constructor(protected scene: Scene, protected positionX: number, protected positionY: number,
                private particleEmitterManager: ParticleEmitterManager, private fillColor = 0x6666ff) {
        super(scene, positionX, positionY);
        this.shipObj.setFillStyle(this.fillColor);
    }

    getPolygon(): Phaser.GameObjects.Polygon {
        const data = [15, 0, 15, 5, 0, 5, 0, 0];
        return this.scene.add.polygon(this.positionX, this.positionY, data, this.fillColor);

    }

    rotate(degrees: number) {
        super.rotate(degrees);
        this.emitter.angle.start += degrees;
        this.emitter.angle.end += degrees;

    }


    moveForward() {
        super.moveForward();

        const newPosition = this.shipObj.getLeftCenter();
        this.emitter.setPosition({
            max: 0,
            min: 0,
            onEmit: undefined,
            random: undefined,
            steps: 0,
            start: newPosition.x, end: newPosition.x}, {
            max: 0,
            min: 0,
            onEmit: undefined,
            random: undefined,
            steps: 0,
            start: newPosition.y, end: newPosition.y});
    }
}