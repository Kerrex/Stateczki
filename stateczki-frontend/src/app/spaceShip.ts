import {SpaceObject} from "./spaceObject";
import Scene = Phaser.Scene;

export class SpaceShip extends SpaceObject {

    getPolygon(): Phaser.GameObjects.Polygon {
        const data = [ 15,0, -5,10, -5, -10 ];
        return this.scene.add.polygon(this.positionX, this.positionY, data, 0x6666ff);
    }

}