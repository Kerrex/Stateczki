import Vector2 = Phaser.Math.Vector2;

export class MathUtils {
    static rotateVector(vectorToRotate: Vector2, degreesToRotate: number): Vector2 {
        const angleDiffInRadian = degreesToRotate * Math.PI / 180;
        const x = Math.cos(angleDiffInRadian) * vectorToRotate.x - Math.sin(angleDiffInRadian) * vectorToRotate.y;
        const y = Math.sin(angleDiffInRadian) * vectorToRotate.x + Math.cos(angleDiffInRadian) * vectorToRotate.y;

        return new Vector2(x, y);
    }
}