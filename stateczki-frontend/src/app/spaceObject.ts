import Scene = Phaser.Scene;
import Polygon = Phaser.GameObjects.Polygon;
import {MathUtils} from "./math-utils";
import Vector2 = Phaser.Math.Vector2;

export abstract class SpaceObject {
    // @ts-ignore
    protected shipObj: Polygon;

    constructor(protected scene: Scene, protected positionX: number, protected positionY: number) {
        this.init();
    }

    protected init() {
        this.shipObj = this.getPolygon();
        this.shipObj.displayOriginX = 0.5;
        this.shipObj.displayOriginY = 0.5;
        this.shipObj.angle = -90;
    }

    abstract getPolygon(): Polygon;

    rotate(degrees: number) {
        this.shipObj.angle += degrees;
    }

    moveForward() {
        const angleToRotate = this.shipObj.angle + 90;
        const movementVector = MathUtils.rotateVector(Vector2.UP, angleToRotate);

        const newX = (this.shipObj.x + 2 * movementVector.x) % window.innerWidth;
        this.shipObj.x = newX < 0 ? window.innerWidth - newX : newX;

        const newY = (this.shipObj.y + 2 * movementVector.y) % window.innerHeight;
        this.shipObj.y = newY < 0 ? window.innerHeight - newY : newY;
    }

}